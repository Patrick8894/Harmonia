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

  // Python Logic Service APIs
  const logicApis = [
    {
      title: 'Hello Logic',
      endpoint: '/api/logic/hello',
      method: 'GET' as const,
      description: 'Simple greeting endpoint from Python logic service',
      language: 'Python' as const,
      parameters: [
        {
          name: 'name',
          type: 'string' as const,
          required: false,
          defaultValue: user || 'World',
          description: 'Name to greet'
        }
      ]
    },
    {
      title: 'Evaluate Expression',
      endpoint: '/api/logic/eval',
      method: 'POST' as const,
      description: 'Evaluate a numeric expression with optional variables',
      language: 'Python' as const,
      bodyParameters: [
        {
          name: 'expression',
          type: 'string' as const,
          required: true,
          placeholder: '2 * x + 3',
          description: 'Mathematical expression to evaluate'
        },
        {
          name: 'variables',
          type: 'object' as const,
          required: false,
          placeholder: '{"x": 5}',
          description: 'Variables for the expression (JSON object)'
        }
      ]
    },
    {
      title: 'Transform Dataset',
      endpoint: '/api/logic/transform',
      method: 'POST' as const,
      description: 'Apply MAP/FILTER/SUM operations on numeric data',
      language: 'Python' as const,
      bodyParameters: [
        {
          name: 'data',
          type: 'array' as const,
          required: true,
          placeholder: '[1, 2, 3, 4, 5]',
          description: 'Input data array'
        },
        {
          name: 'operation',
          type: 'string' as const,
          required: true,
          defaultValue: 'MAP',
          placeholder: 'MAP, FILTER, or SUM',
          description: 'Operation type'
        },
        {
          name: 'expression',
          type: 'string' as const,
          required: false,
          placeholder: 'x * 2',
          description: 'Expression for MAP/FILTER'
        },
        {
          name: 'var_name',
          type: 'string' as const,
          required: false,
          defaultValue: 'x',
          description: 'Variable name in expression'
        }
      ]
    },
    {
      title: 'Plan Tasks',
      endpoint: '/api/logic/plan',
      method: 'POST' as const,
      description: 'Generate a step-by-step plan from a goal',
      language: 'Python' as const,
      bodyParameters: [
        {
          name: 'goal',
          type: 'string' as const,
          required: true,
          placeholder: 'Build a web application',
          description: 'Goal to create a plan for'
        },
        {
          name: 'hints',
          type: 'array' as const,
          required: false,
          placeholder: '["Use React", "Deploy to AWS"]',
          description: 'Optional hints for planning'
        }
      ]
    }
  ];

  // C++ Engine APIs
  const engineApis = [
    {
      title: 'Hello Engine',
      endpoint: '/api/engine/hello',
      method: 'GET' as const,
      description: 'High-performance greeting from C++ compute engine',
      language: 'C++' as const,
      parameters: [
        {
          name: 'name',
          type: 'string' as const,
          required: false,
          defaultValue: user || 'World',
          description: 'Name to greet'
        }
      ]
    },
    {
      title: 'Estimate Pi (Monte Carlo)',
      endpoint: '/api/engine/pi',
      method: 'POST' as const,
      description: 'Estimate π using Monte Carlo simulation',
      language: 'C++' as const,
      bodyParameters: [
        {
          name: 'samples',
          type: 'number' as const,
          required: true,
          defaultValue: 1000000,
          description: 'Number of random samples'
        }
      ]
    },
    {
      title: 'Matrix Multiplication',
      endpoint: '/api/engine/matmul',
      method: 'POST' as const,
      description: 'Multiply two matrices A and B',
      language: 'C++' as const,
      bodyParameters: [
        {
          name: 'A',
          type: 'object' as const,
          required: true,
          placeholder: '{"rows": 2, "cols": 2, "data": [1, 2, 3, 4]}',
          description: 'Matrix A (rows × cols)'
        },
        {
          name: 'B',
          type: 'object' as const,
          required: true,
          placeholder: '{"rows": 2, "cols": 2, "data": [5, 6, 7, 8]}',
          description: 'Matrix B (cols must match A.cols)'
        }
      ]
    },
    {
      title: 'Compute Statistics',
      endpoint: '/api/engine/stats',
      method: 'POST' as const,
      description: 'Calculate mean, variance, stddev, min, max of a dataset',
      language: 'C++' as const,
      bodyParameters: [
        {
          name: 'data',
          type: 'array' as const,
          required: true,
          placeholder: '[1.5, 2.3, 4.7, 3.2, 5.1]',
          description: 'Numeric dataset'
        },
        {
          name: 'sample',
          type: 'boolean' as const,
          required: false,
          defaultValue: true,
          description: 'Use sample variance (n-1) instead of population variance'
        }
      ]
    }
  ];

  // Go Gateway APIs
  const gatewayApis = [
    {
      title: 'Health Check',
      endpoint: '/api/hello',
      method: 'GET' as const,
      description: 'Gateway health check and system status',
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
                  <Grid key={index} size={{ xs: 12, lg: 6 }}>
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
                  <Grid key={index} size={{ xs: 12, lg: 6 }}>
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
                  <Grid key={index} size={{ xs: 12, lg: 6 }}>
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