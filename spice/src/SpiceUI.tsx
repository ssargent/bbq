import * as React from 'react';
import { styled } from '@mui/material/styles';
import MuiAppBar, { AppBarProps } from '@mui/material/AppBar';
import MenuIcon from '@mui/icons-material/Menu';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import NotificationsIcon from '@mui/icons-material/Notifications';
import {
  Alert,
  Box,
  Toolbar,
  IconButton,
  Badge,
  Container,
  Divider,
  List,
  Typography,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Icon,
} from '@mui/material';
import MuiDrawer from '@mui/material/Drawer';
import { Routes, Route, Link } from 'react-router-dom';
import { Simulate } from 'apps/simulate';

const drawerWidth = 240;

interface SpiceAppBarProps extends AppBarProps {
  open?: boolean;
}

const SpiceAppBar = styled(MuiAppBar, {
  shouldForwardProp: (prop) => prop !== 'open',
})<SpiceAppBarProps>(({ theme, open }) => ({
  zIndex: theme.zIndex.drawer + 1,
  transition: theme.transitions.create(['width', 'margin'], {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  ...(open && {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  }),
}));

const Drawer = styled(MuiDrawer, {
  shouldForwardProp: (prop) => prop !== 'open',
})(({ theme, open }) => ({
  '& .MuiDrawer-paper': {
    position: 'relative',
    whiteSpace: 'nowrap',
    width: drawerWidth,
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
    boxSizing: 'border-box',
    ...(!open && {
      overflowX: 'hidden',
      transition: theme.transitions.create('width', {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
      width: theme.spacing(7),
      [theme.breakpoints.up('sm')]: {
        width: theme.spacing(9),
      },
    }),
  },
})) as typeof MuiDrawer;

const SpiceUI = () => {
  const [open, setOpen] = React.useState<boolean>(false);

  const toggleDrawer = () => {
    setOpen(!open);
  };

  return (
    <Box sx={{ display: 'flex' }}>
      <SpiceAppBar position="absolute" open={open}>
        <Toolbar sx={{ pr: '24px' }}>
          <IconButton
            edge="start"
            color="inherit"
            aria-label="open drawer"
            onClick={toggleDrawer}
            sx={{
              marginRight: '36px',
              ...(open && { display: 'none' }),
            }}
          >
            <MenuIcon />
          </IconButton>
          <Typography component="h1" variant="h6" color="inherit" noWrap sx={{ flexGrow: 1 }}>
            BBQd Management Console (Spice)
          </Typography>

          <IconButton color="inherit">
            <Badge badgeContent={4} color="secondary">
              <NotificationsIcon />
            </Badge>
          </IconButton>
        </Toolbar>
      </SpiceAppBar>
      <Drawer variant="permanent" open={open}>
        <Toolbar
          sx={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'flex-end',
            px: [1],
          }}
        >
          <IconButton onClick={toggleDrawer}>
            <ChevronLeftIcon />
          </IconButton>
        </Toolbar>
        <Divider />
        <List component="nav">
          <ListItemButton component={Link} to={'simulate'}>
            <ListItemIcon>
              <Icon>speed_dial</Icon>
            </ListItemIcon>
            <ListItemText primary={'Simulate'} />
          </ListItemButton>{' '}
          <Divider
            sx={{
              my: 1,
            }}
          ></Divider>
        </List>
      </Drawer>
      <Box
        component="main"
        sx={{
          backgroundColor: (theme) =>
            theme.palette.mode == 'light' ? theme.palette.grey[100] : theme.palette.grey[900],
          flexGrow: 1,
          height: '100vh',
          overflow: 'auto',
        }}
      >
        <Toolbar />
        <Container
          maxWidth="lg"
          sx={{
            mt: 4,
            mb: 4,
          }}
        >
          <Routes>
            <Route path="/" element={<SpicePortalHome />} />
            <Route path="/simulate" element={<Simulate />} />
          </Routes>
        </Container>
      </Box>
    </Box>
  );
};

interface SpicePortalHome {}

const SpicePortalHome = () => {
  return (
    <>
      <Alert severity="info">
        <Typography variant="h4">Welcome to the BBQd Console</Typography>
        <Typography variant="body1">
          This portal is the central hub for managing all aspects of the BBQ Platform.
        </Typography>
      </Alert>
    </>
  );
};

export { SpiceUI };
