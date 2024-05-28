import {
  Button,
  FormControl,
  Grid,
  InputLabel,
  OutlinedInput,
  Paper,
  Typography,
} from '@mui/material';
import { Session } from 'bbq/intake/v1/bbq_pb';
import { useState } from 'react';

interface SessionDetailsProps {
  session: Session | undefined;
  onCreateSession: (session: Session) => void;
}

const SessionDetails = ({ session, onCreateSession }: SessionDetailsProps) => {
  const [description, setDescription] = useState<string>(session?.description || '');
  const createSession = () => {
    const newSession = new Session();
    newSession.description = description;
    onCreateSession(newSession);
  };
  return (
    <Paper elevation={1} sx={{ p: 2, m: 2 }}>
      <h2>Session Details</h2>
      <Grid container>
        <Grid item sm={12}>
          <FormControl fullWidth sx={{ m: 1 }}>
            <InputLabel htmlFor="outlined-adornment-amount">Description</InputLabel>
            <OutlinedInput
              id="outlined-adornment-amount"
              label="Description"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
            />
          </FormControl>
          <Typography variant="body2" sx={{ m: 1 }}>
            {session?.id}
          </Typography>
        </Grid>
        <Grid item sm={12} sx={{ textAlign: 'right' }}>
          <Button variant="contained" onClick={() => createSession()}>
            Create Session
          </Button>
        </Grid>
      </Grid>
    </Paper>
  );
};

export { SessionDetails };
