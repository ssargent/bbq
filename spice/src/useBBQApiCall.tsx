import { PromiseClient, Transport, createPromiseClient } from '@connectrpc/connect';
import { IntakeService } from 'bbq/intake/v1/intake_service_connect';
import { useEffect, useState } from 'react';

const getClient = (transport: Transport) => {
  const client = createPromiseClient(IntakeService, transport);
  return client;
};

const useBBQApiCall = (transport: Transport) => {
  const [intakeService, setIntakeService] = useState<
    PromiseClient<typeof IntakeService> | undefined
  >(undefined);

  useEffect(() => {
    if (transport) {
      setIntakeService(getClient(transport));
    }
  }, [transport]);

  return { intakeService };
};

export { useBBQApiCall };
