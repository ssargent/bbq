// cmgrdlg.go: Implements the CentralManagerDelegate interface.  CoreBluetooth
// communicates events asynchronously via callbacks.  This file implements a
// synchronous interface by translating these callbacks into channel
// operations.

package darwin

import (
	"github.com/JuulLabs-OSS/cbgo"
	"github.com/go-ble/ble"
)

func (d *Device) CentralManagerDidUpdateState(cmgr cbgo.CentralManager) {
	d.evl.stateChanged.RxSignal(struct{}{})
}

func (d *Device) DidDiscoverPeripheral(cmgr cbgo.CentralManager, prph cbgo.Peripheral,
	advFields cbgo.AdvFields, rssi int) {

	// The Scan operation is happening in another goroutine. If a scan is still in progress,
	// a chan receive operation on d.advCh will give us a guaranteed-good channel on which
	// we can report this result. If the Scan operation is over, this channel will be closed
	// and we can return early.
	d.connLock.Lock()
	advCh := d.advCh
	d.connLock.Unlock()
	ch := <-advCh
	if ch == nil {
		return
	}

	a := &adv{
		localName: advFields.LocalName,
		rssi:      int(rssi),
		mfgData:   advFields.ManufacturerData,
	}
	if advFields.Connectable != nil {
		a.connectable = *advFields.Connectable
	}
	if advFields.TxPowerLevel != nil {
		a.powerLevel = *advFields.TxPowerLevel
	}
	for _, u := range advFields.ServiceUUIDs {
		a.svcUUIDs = append(a.svcUUIDs, ble.UUID(u))
	}
	for _, sd := range advFields.ServiceData {
		a.svcData = append(a.svcData, ble.ServiceData{
			UUID: ble.UUID(sd.UUID),
			Data: sd.Data,
		})
	}
	a.peerUUID = ble.UUID(prph.Identifier())

	ch <- a
}

func (d *Device) DidConnectPeripheral(cmgr cbgo.CentralManager, prph cbgo.Peripheral) {
	fail := func(err error) {
		d.evl.connected.RxSignal(&eventConnected{
			err: err,
		})
	}

	c, err := newCentralConn(d, prph)
	if err != nil {
		fail(err)
	}

	d.evl.connected.RxSignal(&eventConnected{
		conn: c,
	})
}

func (d *Device) DidDisconnectPeripheral(cmgr cbgo.CentralManager, prph cbgo.Peripheral, err error) {
	c := d.findConn(ble.NewAddr(prph.Identifier().String()))
	if c != nil {
		close(c.done)
	}
}
