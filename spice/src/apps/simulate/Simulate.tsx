import { TransportContext } from 'TransportContext';
import { Reading, SensorReading, Session } from 'bbq/intake/v1/bbq_pb';
import { RecordRequest, SessionRequest } from 'bbq/intake/v1/intake_service_pb';
import { SensorProbe, VirtualBBQ } from 'components';
import { SessionDetails } from 'components/SessionDetails';
import { useContext, useEffect, useRef, useState } from 'react';
import { useBBQApiCall } from 'useBBQApiCall';

const Simulate = () => {
  const transport = useContext(TransportContext);
  const { intakeService } = useBBQApiCall(transport);
  const [session, setSession] = useState<Session | undefined>(undefined);
  const [error, setError] = useState<string | undefined>(undefined);
  const [probes, setProbes] = useState<SensorProbe[]>([
    { number: 0, maxTemp: 1200, minTemp: 100 },
    { number: 1, maxTemp: 1200, minTemp: 100 },
    { number: 2, maxTemp: 1200, minTemp: 100 },
    { number: 3, maxTemp: 1200, minTemp: 100 },
  ]);
  const [intervalHandle, setIntervalHandle] = useState<number | undefined>(undefined);
  const [cooking, setCooking] = useState<boolean>(false);
  const [currentTemps, setCurrentTemps] = useState<number[]>([125, 125, 125, 125]);
  const createSession = (session: Session) => {
    if (!intakeService) throw new Error('Intake service not available');

    const request = new SessionRequest();
    request.description = session.description;

    intakeService
      .session(request)
      .then((resp) => {
        setSession(resp.session);
      })
      .catch((err) => {
        console.error(err);
        setError(err.message);
      });
  };

  useInterval(() => recordTemperatures(), cooking == true ? 1500 : null);

  const recordTemperatures = () => {
    if (!intakeService) throw new Error('Intake service not available');
    if (!session) throw new Error('Session not available');

    const request = new RecordRequest();

    const reading0 = new SensorReading();
    reading0.sensorNumber = 0;
    reading0.temperature = currentTemps[0];

    const reading1 = new SensorReading();
    reading1.sensorNumber = 1;
    reading1.temperature = currentTemps[1];

    const reading2 = new SensorReading();
    reading2.sensorNumber = 2;
    reading2.temperature = currentTemps[2];

    const reading3 = new SensorReading();
    reading3.sensorNumber = 3;
    reading3.temperature = currentTemps[3];

    const reading = new Reading();
    // add our temp probe readings to the reading record to be recorded
    reading.readings = [reading0, reading1, reading2, reading3];
    reading.sessionId = session.id || '';

    // this api allows sending multiple readings at once, in a batch - but we're doing one at a time.
    request.reading = [reading];

    intakeService
      .record(request)
      .then((resp) => {
        console.log(resp);
      })
      .catch((err) => {
        console.error(err);
        setError(err.message);
      });
  };

  const startCooking = () => {
    setCooking(true); // start the interval
  };

  const stopCooking = () => {
    setCooking(false); // stop the interval
  };

  useEffect(() => {}, []);

  return (
    <>
      <h1>Simulate</h1>
      <SessionDetails
        session={session}
        onCreateSession={(session: Session) => {
          createSession(session);
        }}
      />
      <VirtualBBQ
        probes={probes}
        session={session}
        onTemperatureChange={setCurrentTemps}
        startCooking={startCooking}
        stopCooking={stopCooking}
        cooking={cooking}
      />

      {currentTemps.map((temp, index) => (
        <div key={index}>
          Probe {index}: {temp}°F
        </div>
      ))}
    </>
  );
};

// Code adapted from https://overreacted.io/making-setinterval-declarative-with-react-hooks/
const useInterval = (callback: () => void, delay: number | null) => {
  const savedCallback = useRef<() => void>();

  useEffect(() => {
    savedCallback.current = callback;
  }, [callback]);

  useEffect(() => {
    const tick = () => {
      if (savedCallback.current) {
        savedCallback.current();
      }
    };

    if (delay !== null) {
      const id = setInterval(tick, delay);
      return () => clearInterval(id);
    }
  }, [delay]);
};

export { Simulate };
