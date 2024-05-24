import { Session } from 'bbq/intake/v1/bbq_pb';
import { SensorProbe, VirtualBBQ } from 'components';
import { SessionDetails } from 'components/SessionDetails';
import { useEffect, useState } from 'react';

const Simulate = () => {
  const [session, setSession] = useState<Session | undefined>(undefined);
  const [probes, setProbes] = useState<SensorProbe[]>([
    { number: 0, maxTemp: 1200, minTemp: 100 },
    { number: 1, maxTemp: 1200, minTemp: 100 },
    { number: 2, maxTemp: 1200, minTemp: 100 },
    { number: 3, maxTemp: 1200, minTemp: 100 },
  ]);

  useEffect(() => {}, []);

  return (
    <>
      <h1>Simulate</h1>
      <SessionDetails
        session={session}
        onCreateSession={(session: Session) => {
          setSession(session);
        }}
      />
      <VirtualBBQ probes={probes} session={session} />
    </>
  );
};

export { Simulate };
