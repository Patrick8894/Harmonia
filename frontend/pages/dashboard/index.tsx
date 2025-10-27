import Header from '../../components/Header';
import { useAuth } from '../../contexts/AuthContext';
import ApiEndpointCard from '../../components/ApiEndpointCard';
import {
  Box,
  Container,
  Typography,
  Paper,
  Stack,
  Tabs,
  Tab,
  Chip,
  Grid
} from '@mui/material';
import { useState } from 'react';
import {
  Memory as MemoryIcon,
  Psychology as PsychologyIcon,
  Api as ApiIcon
} from '@mui/icons-material';

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

function TabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;
  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`tabpanel-${index}`}
      aria-labelledby={`tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ py: 3 }}>{children}</Box>}
    </div>
  );
}

export default function Dashboard() {
  const { user, loading } = useAuth();
  const [tabValue, setTabValue] = useState(0);

  if (loading) return null;

  const handleTabChange = (event: React.SyntheticEvent, newValue: number) => {
    setTabValue(newValue);
  };

  // Define your API endpoints here
  const logicApis = [
    {
      title: 'Hello Logic',
      endpoint: '/api/logic/hello',
      description: 'Simple greeting endpoint from Python logic service',
      language: 'Python' as const,
      parameters: [
        {
          name: 'name',
          type: 'string',
          required: false,
          defaultValue: user || 'Patrick',
          description: 'Name to greet'
        }
      ]
    },
    {
      title: 'Data Processing',
      endpoint: '/api/logic/process',
      description: 'Process and transform data using Python workflows',
      language: 'Python' as const,
      parameters: [
        {
          name: 'data',
          type: 'string',
          required: true,
          description: 'Input data to process'
        },
        {
          name: 'operation',
          type: 'string',
          required: false,
          defaultValue: 'transform',
          description: 'Operation type (transform, validate, analyze)'
        }
      ]
    }
  ];

  const engineApis = [
    {
      title: 'Hello Engine',
      endpoint: '/api/engine/hello',
      description: 'High-performance greeting from C++ compute engine',
      language: 'C++' as const,
      parameters: [
        {
          name: 'name',
          type: 'string',
          required: false,
          defaultValue: user || 'Patrick',
          description: 'Name to greet'
        }
      ]
    },
    {
      title: 'Matrix Computation',
      endpoint: '/api/engine/compute',
      description: 'Perform intensive numerical computations',
      language: 'C++' as const,
      parameters: [
        {
          name: 'size',
          type: 'number',
          required: true,
          defaultValue: '100',
          description: 'Matrix size (NxN)'
        },
        {
          name: 'iterations',
          type: 'number',
          required: false,
          defaultValue: '1000',
          description: 'Number of iterations'
        }
      ]
    }
  ];

  const gatewayApis = [
    {
      title: 'Health Check',
      endpoint: '/api/hello',
      description: 'Gateway health check and system status',
      language: 'Go' as const
    },
    {
      title: 'System Metrics',
      endpoint: '/api/metrics',
      description: 'Real-time system performance metrics',
      language: 'Go' as const
    }
  ];

  return (
    <>
      <Header user={user} />
      <Box
        sx={{
          background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
          minHeight: '30vh',
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
            <Typography variant="h3" component="h1" fontWeight={700}>
              Dashboard
            </Typography>
            <Typography variant="h6" sx={{ opacity: 0.95 }}>
              Welcome back, <strong>{user}</strong> 🎉
            </Typography>
            <Typography variant="body1" sx={{ opacity: 0.9, maxWidth: 700 }}>
              Test and interact with the hybrid compute services. Select a service category below to explore available endpoints.
            </Typography>
          </Stack>
        </Container>
      </Box>

      <Container maxWidth="lg" sx={{ mt: -6, mb: 6, position: 'relative', zIndex: 2 }}>
        <Paper elevation={4} sx={{ borderRadius: 2, overflow: 'hidden' }}>
          <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
            <Tabs
              value={tabValue}
              onChange={handleTabChange}
              aria-label="API service tabs"
              sx={{
                px: 2,
                '& .MuiTab-root': {
                  textTransform: 'none',
                  fontWeight: 600,
                  fontSize: '1rem'
                }
              }}
            >
              <Tab
                label={
                  <Stack direction="row" spacing={1} alignItems="center">
                    <PsychologyIcon />
                    <span>Logic Service</span>
                    <Chip label={logicApis.length} size="small" />
                  </Stack>
                }
              />
              <Tab
                label={
                  <Stack direction="row" spacing={1} alignItems="center">
                    <MemoryIcon />
                    <span>Compute Engine</span>
                    <Chip label={engineApis.length} size="small" />
                  </Stack>
                }
              />
              <Tab
                label={
                  <Stack direction="row" spacing={1} alignItems="center">
                    <ApiIcon />
                    <span>Gateway</span>
                    <Chip label={gatewayApis.length} size="small" />
                  </Stack>
                }
              />
            </Tabs>
          </Box>

          <Box sx={{ p: 3 }}>
            <TabPanel value={tabValue} index={0}>
              <Stack spacing={2} mb={3}>
                <Typography variant="h5" fontWeight={600}>
                  Python Logic Service
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Business logic and data processing workflows powered by Python
                </Typography>
              </Stack>
              <Grid container spacing={3}>
                {logicApis.map((api, index) => (
                  <Grid key={index} size={{ xs: 12, md: 6 }}>
                    <ApiEndpointCard {...api} />
                  </Grid>
                ))}
              </Grid>
            </TabPanel>

            <TabPanel value={tabValue} index={1}>
              <Stack spacing={2} mb={3}>
                <Typography variant="h5" fontWeight={600}>
                  C++ Compute Engine
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  High-performance numerical computations and intensive processing
                </Typography>
              </Stack>
              <Grid container spacing={3}>
                {engineApis.map((api, index) => (
                  <Grid key={index} size={{ xs: 12, md: 6 }}>
                    <ApiEndpointCard {...api} />
                  </Grid>
                ))}
              </Grid>
            </TabPanel>

            <TabPanel value={tabValue} index={2}>
              <Stack spacing={2} mb={3}>
                <Typography variant="h5" fontWeight={600}>
                  Go API Gateway
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Core orchestration and routing services
                </Typography>
              </Stack>
              <Grid container spacing={3}>
                {gatewayApis.map((api, index) => (
                  <Grid key={index} size={{ xs: 12, md: 6 }}>
                    <ApiEndpointCard {...api} />
                  </Grid>
                ))}
              </Grid>
            </TabPanel>
          </Box>
        </Paper>
      </Container>
    </>
  );
}