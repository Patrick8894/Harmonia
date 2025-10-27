import Link from 'next/link';
import { AppBar, Toolbar, Typography, Button, Box, Container, Chip, Menu, MenuItem } from '@mui/material';
import { 
  Home as HomeIcon, 
  Dashboard as DashboardIcon, 
  Description as DocsIcon,
  Login as LoginIcon,
  Logout as LogoutIcon,
  Person as PersonIcon
} from '@mui/icons-material';
import { useState } from 'react';
import { logout } from '../lib/api';
import { useAuth } from '../contexts/AuthContext';

export default function Header({ user }: { user?: string }) {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const { refresh } = useAuth();
  const open = Boolean(anchorEl);

  const handleClick = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleLogout = async () => {
    await logout();
    await refresh();
    handleClose();
  };
  return (
    <AppBar 
      position="static" 
      elevation={0}
      sx={{ 
        bgcolor: 'white', 
        borderBottom: '1px solid',
        borderColor: 'divider'
      }}
    >
      <Container maxWidth="lg">
        <Toolbar sx={{ px: { xs: 0 }, justifyContent: 'space-between' }}>
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 3 }}>
            <Link href="/" passHref legacyBehavior>
              <Typography 
                variant="h6" 
                component="a" 
                sx={{ 
                  fontWeight: 700,
                  background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
                  backgroundClip: 'text',
                  WebkitBackgroundClip: 'text',
                  WebkitTextFillColor: 'transparent',
                  letterSpacing: '-0.5px',
                  textDecoration: 'none',
                  cursor: 'pointer',
                  '&:hover': {
                    opacity: 0.8
                  }
                }}
              >
                Harmonia
              </Typography>
            </Link>
            <Box sx={{ display: { xs: 'none', md: 'flex' }, gap: 1 }}>
              <Link href="/" passHref legacyBehavior>
                <Button 
                  component="a" 
                  startIcon={<HomeIcon />}
                  sx={{ 
                    color: 'text.primary',
                    textTransform: 'none',
                    fontWeight: 500
                  }}
                >
                  Home
                </Button>
              </Link>
              <Link href="/dashboard" passHref legacyBehavior>
                <Button 
                  component="a" 
                  startIcon={<DashboardIcon />}
                  sx={{ 
                    color: 'text.primary',
                    textTransform: 'none',
                    fontWeight: 500
                  }}
                >
                  Dashboard
                </Button>
              </Link>
              <Link href="/swagger/index.html" target="_blank" rel="noreferrer" passHref legacyBehavior>
                <Button 
                  component="a" 
                  startIcon={<DocsIcon />}
                  sx={{ 
                    color: 'text.primary',
                    textTransform: 'none',
                    fontWeight: 500
                  }}
                >
                  API Docs
                </Button>
              </Link>
            </Box>
          </Box>
          
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
            {user ? (
              <>
                <Chip 
                  label={user}
                  icon={<PersonIcon />}
                  size="small"
                  onClick={handleClick}
                  sx={{ 
                    fontWeight: 600,
                    cursor: 'pointer',
                    background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
                    color: 'white',
                    '& .MuiChip-icon': {
                      color: 'white'
                    },
                    '&:hover': {
                      background: 'linear-gradient(135deg, #5568d3 0%, #63428d 100%)',
                    }
                  }}
                />
                <Menu
                  anchorEl={anchorEl}
                  open={open}
                  onClose={handleClose}
                  anchorOrigin={{
                    vertical: 'bottom',
                    horizontal: 'right',
                  }}
                  transformOrigin={{
                    vertical: 'top',
                    horizontal: 'right',
                  }}
                >
                  <MenuItem onClick={handleLogout}>
                    <LogoutIcon sx={{ mr: 1, fontSize: 20 }} />
                    Logout
                  </MenuItem>
                </Menu>
              </>
            ) : (
              <Link href="/login" passHref legacyBehavior>
                <Button 
                  component="a" 
                  variant="outlined"
                  startIcon={<LoginIcon />}
                  sx={{ 
                    textTransform: 'none',
                    fontWeight: 600,
                    borderColor: '#667eea',
                    color: '#667eea',
                    '&:hover': {
                      borderColor: '#5568d3',
                      color: '#5568d3',
                      backgroundColor: 'rgba(102, 126, 234, 0.04)'
                    }
                  }}
                >
                  Login
                </Button>
              </Link>
            )}
          </Box>
        </Toolbar>
      </Container>
    </AppBar>
  );
}