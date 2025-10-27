import { useState } from 'react';
import {
  Card,
  CardContent,
  Typography,
  Button,
  Stack,
  Chip,
  TextField,
  Box,
  Collapse,
  Alert
} from '@mui/material';
import {
  Launch as LaunchIcon,
  ExpandMore as ExpandMoreIcon,
  Send as SendIcon
} from '@mui/icons-material';

interface ApiParameter {
  name: string;
  type: string;
  required: boolean;
  defaultValue?: string;
  description?: string;
}

interface ApiEndpointCardProps {
  title: string;
  endpoint: string;
  method?: 'GET' | 'POST' | 'PUT' | 'DELETE';
  description: string;
  language: 'Go' | 'Python' | 'C++';
  parameters?: ApiParameter[];
}

const languageColors = {
  'Go': '#00ADD8',
  'Python': '#3776AB',
  'C++': '#00599C'
};

export default function ApiEndpointCard({
  title,
  endpoint,
  method = 'GET',
  description,
  language,
  parameters = []
}: ApiEndpointCardProps) {
  const [expanded, setExpanded] = useState(false);
  const [paramValues, setParamValues] = useState<Record<string, string>>(
    parameters.reduce((acc, param) => ({
      ...acc,
      [param.name]: param.defaultValue || ''
    }), {})
  );
  const [response, setResponse] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const buildUrl = () => {
    if (parameters.length === 0) return endpoint;
    
    const queryParams = new URLSearchParams();
    Object.entries(paramValues).forEach(([key, value]) => {
      if (value) queryParams.append(key, value);
    });
    
    const queryString = queryParams.toString();
    return queryString ? `${endpoint}?${queryString}` : endpoint;
  };

  const handleTest = async () => {
    setLoading(true);
    setError(null);
    setResponse(null);

    try {
      const url = buildUrl();
      const res = await fetch(url, { method });
      const data = await res.text();
      setResponse(data);
    } catch (e: any) {
      setError(e?.message || 'Request failed');
    } finally {
      setLoading(false);
    }
  };

  const hasParameters = parameters.length > 0;

  return (
    <Card variant="outlined" sx={{ '&:hover': { boxShadow: 2 }, transition: 'box-shadow 0.2s' }}>
      <CardContent>
        <Stack spacing={2}>
          <Box>
            <Stack direction="row" spacing={1} alignItems="center" mb={1}>
              <Typography variant="h6" fontWeight={600}>
                {title}
              </Typography>
              <Chip
                label={language}
                size="small"
                sx={{
                  bgcolor: languageColors[language],
                  color: 'white',
                  fontWeight: 600,
                  fontSize: '0.7rem'
                }}
              />
              <Chip
                label={method}
                size="small"
                variant="outlined"
                sx={{ fontWeight: 600, fontSize: '0.7rem' }}
              />
            </Stack>
            <Typography
              variant="body2"
              fontFamily="monospace"
              sx={{
                bgcolor: 'grey.100',
                p: 1,
                borderRadius: 1,
                mb: 1,
                wordBreak: 'break-all'
              }}
            >
              {endpoint}
            </Typography>
            <Typography variant="body2" color="text.secondary">
              {description}
            </Typography>
          </Box>

          <Stack direction="row" spacing={1}>
            {hasParameters ? (
              <>
                <Button
                  variant="outlined"
                  size="small"
                  onClick={() => setExpanded(!expanded)}
                  endIcon={
                    <ExpandMoreIcon
                      sx={{
                        transform: expanded ? 'rotate(180deg)' : 'rotate(0deg)',
                        transition: 'transform 0.2s'
                      }}
                    />
                  }
                  sx={{ textTransform: 'none' }}
                >
                  {expanded ? 'Hide' : 'Show'} Parameters
                </Button>
                <Button
                  variant="contained"
                  size="small"
                  endIcon={<SendIcon />}
                  onClick={handleTest}
                  disabled={loading}
                  sx={{
                    textTransform: 'none',
                    background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
                    '&:hover': {
                      background: 'linear-gradient(135deg, #5568d3 0%, #63428d 100%)',
                    }
                  }}
                >
                  {loading ? 'Testing...' : 'Test API'}
                </Button>
              </>
            ) : (
              <Button
                variant="outlined"
                size="small"
                endIcon={<LaunchIcon />}
                href={buildUrl()}
                target="_blank"
                rel="noreferrer"
                sx={{ textTransform: 'none' }}
              >
                Open in New Tab
              </Button>
            )}
          </Stack>

          {hasParameters && (
            <Collapse in={expanded}>
              <Stack spacing={2} sx={{ pt: 1 }}>
                {parameters.map((param) => (
                  <TextField
                    key={param.name}
                    label={param.name}
                    value={paramValues[param.name] || ''}
                    onChange={(e) =>
                      setParamValues((prev) => ({
                        ...prev,
                        [param.name]: e.target.value
                      }))
                    }
                    size="small"
                    required={param.required}
                    helperText={param.description}
                    placeholder={param.defaultValue}
                  />
                ))}
              </Stack>
            </Collapse>
          )}

          {error && (
            <Alert severity="error" onClose={() => setError(null)}>
              {error}
            </Alert>
          )}

          {response && (
            <Box
              sx={{
                bgcolor: 'grey.100',
                p: 2,
                borderRadius: 1,
                maxHeight: 200,
                overflow: 'auto'
              }}
            >
              <Typography variant="caption" color="text.secondary" fontWeight={600} display="block" mb={1}>
                Response:
              </Typography>
              <Typography variant="body2" fontFamily="monospace" sx={{ whiteSpace: 'pre-wrap' }}>
                {response}
              </Typography>
            </Box>
          )}
        </Stack>
      </CardContent>
    </Card>
  );
}