import { useAuth } from '../contexts/AuthContext';
import Header from '../components/Header';
import Link from "next/link";
import { 
  Box, 
  Button, 
  Typography, 
  Paper, 
  List, 
  ListItem, 
  Divider, 
  Stack,
  Chip,
  Card,
  CardContent,
  Container,
  Grid
} from '@mui/material';
import { 
  Code as CodeIcon, 
  Speed as SpeedIcon, 
  AccountTree as TreeIcon,
  Launch as LaunchIcon 
} from '@mui/icons-material';

export default function Home() {
  const { user } = useAuth();

  const apiEndpoints = [
    {
      path: '/api/hello',
      description: 'Simple health check endpoint',
      language: 'Go',
      color: '#00ADD8'
    },
    {
      path: `/api/logic/hello?name=${encodeURIComponent(user || "Patrick")}`,
      description: 'Business logic service with Python',
      language: 'Python',
      color: '#3776AB'
    },
    {
      path: `/api/engine/hello?name=${encodeURIComponent(user || "Patrick")}`,
      description: 'High-performance compute engine in C++',
      language: 'C++',
      color: '#00599C'
    }
  ];

  return (
    <>
      <Header user={user} />
      <Box
        sx={{
          background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
          minHeight: '40vh',
          display: 'flex',
          alignItems: 'center',
          color: 'white',
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
        <Container maxWidth="lg" sx={{ position: 'relative', zIndex: 1 }}>
          <Stack spacing={2}>
            <Typography variant="h2" component="h1" fontWeight={700}>
              Harmonia
            </Typography>
            <Typography variant="h5" sx={{ opacity: 0.95, maxWidth: 600 }}>
              Hybrid Compute Platform
            </Typography>
            <Typography variant="body1" sx={{ opacity: 0.9, maxWidth: 700 }}>
              Orchestrating Go, Python, and C++ services through seamless cross-language RPC integration
            </Typography>
            <Stack direction="row" spacing={2} sx={{ mt: 3 }}>
              <Chip 
                icon={<CodeIcon />} 
                label="Go" 
                sx={{ bgcolor: 'rgba(255,255,255,0.2)', color: 'white', fontWeight: 600 }} 
              />
              <Chip 
                icon={<CodeIcon />} 
                label="Python" 
                sx={{ bgcolor: 'rgba(255,255,255,0.2)', color: 'white', fontWeight: 600 }} 
              />
              <Chip 
                icon={<CodeIcon />} 
                label="C++" 
                sx={{ bgcolor: 'rgba(255,255,255,0.2)', color: 'white', fontWeight: 600 }} 
              />
            </Stack>
          </Stack>
        </Container>
      </Box>

      <Container maxWidth="lg" sx={{ mt: -8, mb: 6, position: 'relative', zIndex: 2 }}>
        <Paper elevation={4} sx={{ p: 4, borderRadius: 2 }}>

          <Typography variant="h5" gutterBottom fontWeight={600} mb={3}>
            System Architecture
          </Typography>

          <Grid container spacing={3} mb={4}>
            <Grid size={{ xs: 12, md: 4 }}>
              <Card variant="outlined" sx={{ height: '100%' }}>
                <CardContent>
                  <Stack spacing={1}>
                    <TreeIcon sx={{ fontSize: 40, color: '#667eea' }} />
                    <Typography variant="h6" fontWeight={600}>
                      Orchestration
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                      Go-based API gateway managing cross-service communication and request routing
                    </Typography>
                  </Stack>
                </CardContent>
              </Card>
            </Grid>
            <Grid size={{ xs: 12, md: 4 }}>
              <Card variant="outlined" sx={{ height: '100%' }}>
                <CardContent>
                  <Stack spacing={1}>
                    <CodeIcon sx={{ fontSize: 40, color: '#6e5bb7' }} />
                    <Typography variant="h6" fontWeight={600}>
                      Business Logic
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                      Python services handling complex workflows and data processing pipelines
                    </Typography>
                  </Stack>
                </CardContent>
              </Card>
            </Grid>
            <Grid size={{ xs: 12, md: 4 }}>
              <Card variant="outlined" sx={{ height: '100%' }}>
                <CardContent>
                  <Stack spacing={1}>
                    <SpeedIcon sx={{ fontSize: 40, color: '#764ba2' }} />
                    <Typography variant="h6" fontWeight={600}>
                      Compute Engine
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                      C++ backend delivering high-performance numerical computations
                    </Typography>
                  </Stack>
                </CardContent>
              </Card>
            </Grid>
          </Grid>

          <Divider sx={{ my: 4 }} />

          <Typography variant="h5" gutterBottom fontWeight={600} mb={2}>
            API Endpoints
          </Typography>
          <Typography variant="body2" color="text.secondary" mb={3}>
            Test the hybrid compute system with these live endpoints
          </Typography>

          <List sx={{ p: 0 }}>
            {apiEndpoints.map((endpoint, index) => (
              <ListItem 
                key={index}
                sx={{ 
                  px: 0,
                  py: 2,
                  borderBottom: index < apiEndpoints.length - 1 ? '1px solid' : 'none',
                  borderColor: 'divider'
                }}
              >
                <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2} sx={{ width: '100%' }} alignItems={{ sm: 'center' }}>
                  <Box sx={{ flex: 1 }}>
                    <Stack direction="row" spacing={1} alignItems="center" mb={0.5}>
                      <Typography variant="body1" fontWeight={600} fontFamily="monospace">
                        {endpoint.path.split('?')[0]}
                      </Typography>
                      <Chip 
                        label={endpoint.language} 
                        size="small" 
                        sx={{ 
                          bgcolor: endpoint.color, 
                          color: 'white',
                          fontWeight: 600,
                          fontSize: '0.7rem'
                        }} 
                      />
                    </Stack>
                    <Typography variant="body2" color="text.secondary">
                      {endpoint.description}
                    </Typography>
                  </Box>
                  <Link href={endpoint.path} target="_blank" rel="noreferrer" passHref legacyBehavior>
                    <Button 
                      component="a" 
                      variant="outlined" 
                      endIcon={<LaunchIcon />}
                      size="small"
                      sx={{ textTransform: 'none', whiteSpace: 'nowrap' }}
                    >
                      Test
                    </Button>
                  </Link>
                </Stack>
              </ListItem>
            ))}
          </List>
        </Paper>
      </Container>
    </>
  );
}