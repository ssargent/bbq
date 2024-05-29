//@ts-check
import { createTheme } from '@mui/material';
import { deepPurple, lightGreen, red } from '@mui/material/colors';

const theme = createTheme({
  palette: {
    primary: { main: deepPurple[900] },
    secondary: { main: lightGreen[500] },
    error: {
      main: red.A700,
    },
  },
});

export default theme;
