import { FormEvent, useState, useEffect } from 'react';
import { useRouter } from 'next/router';
import { register } from '../lib/api';
import { useAuth } from '../contexts/AuthContext';
import Header from '../components/Header';
import Link from 'next/link';
import {
  Box,
  Container,
  Paper,
  TextField,
  Button,
  Typography,
  Alert,
  Stack,
  CircularProgress
} from '@mui/material';
import { PersonAdd as RegisterIcon } from '@mui/icons-material';

export default function RegisterPage() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [err, setErr] = useState('');
  const [loading, setLoading] = useState(false);
  const router = useRouter();
  const { user, refresh } = useAuth();

  // If already logged in, redirect to home
  useEffect(() => {
    if (user) {
      router.replace('/');
    }
  }, [user, router]);

  const onSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setErr('');

    // Validation
    if (password !== confirmPassword) {
      setErr('Passwords do not match');
      return;
    }

    if (password.length < 6) {
      setErr('Password must be at least 6 characters');
      return;
    }

    setLoading(true);
    try {
      await register(username, password);
      await refresh();
      router.replace('/');
    } catch (e: any) {
      setErr(e?.message || 'Registration failed');
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <Header user={user} />
      <Box
        sx={{
          minHeight: 'calc(100vh - 64px)',
          background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
          display: 'flex',
          alignItems: 'center',
          position: 'relative',
          overflow: 'hidden',
          '&::before': {
            content: '""',
            position: 'absolute',
            top: 0,
            left: 0,
            right: 0,
            bottom: 0,
            background: 'url("data:image/svg+xml,%3Csvg width="60" height="60" viewBox="0 0 60 60" xmlns="http://www.w3.org/2000/svg"%3E%3Cg fill="none" fill-rule="evenodd"%3E%3Cg fill="%23ffffff" fill-opacity="0.05"%3E%3Cpath d="M36 34v-4h-2v4h-4v2h4v4h2v-4h4v-2h-4zm0-30V0h-2v4h-4v2h4v4h2V6h4V4h-4zM6 34v-4H4v4H0v2h4v4h2v-4h4v-2H6zM6 4V0H4v4H0v2h4v4h2V6h4V4H6z"/%3E%3C/g%3E%3C/g%3E%3C/svg%3E")',
          }
        }}
      >
        <Container maxWidth="sm" sx={{ position: 'relative', zIndex: 1 }}>
          <Paper elevation={4} sx={{ p: 5, borderRadius: 2 }}>
            <Stack spacing={3}>
              <Box sx={{ textAlign: 'center' }}>
                <Box
                  sx={{
                    width: 64,
                    height: 64,
                    borderRadius: '50%',
                    background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                    margin: '0 auto 16px',
                  }}
                >
                  <RegisterIcon sx={{ fontSize: 32, color: 'white' }} />
                </Box>
                <Typography variant="h4" fontWeight={700} gutterBottom>
                  Create Account
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Join the Harmonia hybrid compute platform
                </Typography>
              </Box>

              {err && (
                <Alert severity="error" onClose={() => setErr('')}>
                  {err}
                </Alert>
              )}

              <form onSubmit={onSubmit}>
                <Stack spacing={3}>
                  <TextField
                    label="Username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    fullWidth
                    required
                    autoComplete="username"
                    autoFocus
                    disabled={loading}
                  />
                  <TextField
                    label="Password"
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    fullWidth
                    required
                    autoComplete="new-password"
                    disabled={loading}
                    helperText="Minimum 6 characters"
                  />
                  <TextField
                    label="Confirm Password"
                    type="password"
                    value={confirmPassword}
                    onChange={(e) => setConfirmPassword(e.target.value)}
                    fullWidth
                    required
                    autoComplete="new-password"
                    disabled={loading}
                  />
                  <Button
                    type="submit"
                    variant="contained"
                    size="large"
                    fullWidth
                    disabled={loading}
                    sx={{
                      textTransform: 'none',
                      fontWeight: 600,
                      py: 1.5,
                      background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
                      '&:hover': {
                        background: 'linear-gradient(135deg, #5568d3 0%, #63428d 100%)',
                      }
                    }}
                  >
                    {loading ? (
                      <>
                        <CircularProgress size={20} sx={{ mr: 1, color: 'white' }} />
                        Creating account...
                      </>
                    ) : (
                      'Create Account'
                    )}
                  </Button>
                </Stack>
              </form>

              <Box sx={{ textAlign: 'center', pt: 2, borderTop: '1px solid', borderColor: 'divider' }}>
                <Typography variant="body2" color="text.secondary">
                  Already have an account?{' '}
                  <Link href="/login" passHref legacyBehavior>
                    <Typography
                      component="a"
                      variant="body2"
                      sx={{
                        color: '#667eea',
                        textDecoration: 'none',
                        fontWeight: 600,
                        '&:hover': {
                          textDecoration: 'underline',
                        }
                      }}
                    >
                      Sign In
                    </Typography>
                  </Link>
                </Typography>
              </Box>
            </Stack>
          </Paper>
        </Container>
      </Box>
    </>
  );
}