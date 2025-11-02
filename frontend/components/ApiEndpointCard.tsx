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
  Alert,
  FormControlLabel,
  Switch
} from '@mui/material';
import {
  Launch as LaunchIcon,
  ExpandMore as ExpandMoreIcon,
  Send as SendIcon
} from '@mui/icons-material';

interface ApiParameter {
  name: string;
  type: 'string' | 'number' | 'boolean' | 'array' | 'object';
  required: boolean;
  defaultValue?: any;
  description?: string;
  placeholder?: string;
}

interface ApiEndpointCardProps {
  title: string;
  endpoint: string;
  method?: 'GET' | 'POST' | 'PUT' | 'DELETE';
  description: string;
  language: 'Go' | 'Python' | 'C++';
  parameters?: ApiParameter[];
  bodyParameters?: ApiParameter[];
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
  parameters = [],
  bodyParameters = []
}: ApiEndpointCardProps) {
  const [expanded, setExpanded] = useState(false);
  const [queryParams, setQueryParams] = useState<Record<string, string>>(
    parameters.reduce((acc, param) => ({
      ...acc,
      [param.name]: param.defaultValue || ''
    }), {})
  );
  const [bodyParams, setBodyParams] = useState<Record<string, any>>(
    bodyParameters.reduce((acc, param) => ({
      ...acc,
      [param.name]: param.defaultValue !== undefined ? param.defaultValue : ''
    }), {})
  );
  const [response, setResponse] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const buildUrl = () => {
    if (parameters.length === 0) return endpoint;
    
    const queryParamsObj = new URLSearchParams();
    Object.entries(queryParams).forEach(([key, value]) => {
      if (value) queryParamsObj.append(key, value);
    });
    
    const queryString = queryParamsObj.toString();
    return queryString ? `${endpoint}?${queryString}` : endpoint;
  };

  const buildRequestBody = () => {
    const body: Record<string, any> = {};
    
    bodyParameters.forEach((param) => {
      const value = bodyParams[param.name];
      
      if (value === '' || value === null || value === undefined) {
        if (param.required) {
          body[param.name] = param.defaultValue || '';
        }
        return;
      }

      // Type conversion
      if (param.type === 'number') {
        body[param.name] = parseFloat(value);
      } else if (param.type === 'boolean') {
        body[param.name] = value;
      } else if (param.type === 'array') {
        try {
          body[param.name] = typeof value === 'string' ? JSON.parse(value) : value;
        } catch {
          body[param.name] = [];
        }
      } else if (param.type === 'object') {
        try {
          body[param.name] = typeof value === 'string' ? JSON.parse(value) : value;
        } catch {
          body[param.name] = {};
        }
      } else {
        body[param.name] = value;
      }
    });

    return body;
  };

  const handleTest = async () => {
    setLoading(true);
    setError(null);
    setResponse(null);

    try {
      const url = buildUrl();
      const options: RequestInit = {
        method,
        headers: method !== 'GET' ? { 'Content-Type': 'application/json' } : undefined,
        body: method !== 'GET' ? JSON.stringify(buildRequestBody()) : undefined
      };

      const res = await fetch(url, options);
      const contentType = res.headers.get('content-type');
      
      let data;
      if (contentType?.includes('application/json')) {
        data = await res.json();
        setResponse(JSON.stringify(data, null, 2));
      } else {
        data = await res.text();
        setResponse(data);
      }

      if (!res.ok) {
        setError(`Request failed with status ${res.status}`);
      }
    } catch (e: any) {
      setError(e?.message || 'Request failed');
    } finally {
      setLoading(false);
    }
  };

  const hasParameters = parameters.length > 0 || bodyParameters.length > 0;
  const canOpenDirectly = method === 'GET' && bodyParameters.length === 0;

  const renderParameter = (param: ApiParameter, value: any, onChange: (val: any) => void) => {
    if (param.type === 'boolean') {
      return (
        <FormControlLabel
          key={param.name}
          control={
            <Switch
              checked={value || false}
              onChange={(e) => onChange(e.target.checked)}
            />
          }
          label={
            <Box>
              <Typography variant="body2" fontWeight={600}>
                {param.name} {param.required && <span style={{ color: 'red' }}>*</span>}
              </Typography>
              {param.description && (
                <Typography variant="caption" color="text.secondary">
                  {param.description}
                </Typography>
              )}
            </Box>
          }
        />
      );
    }

    return (
      <TextField
        key={param.name}
        label={param.name}
        value={value || ''}
        onChange={(e) => onChange(e.target.value)}
        size="small"
        fullWidth
        required={param.required}
        helperText={param.description}
        placeholder={param.placeholder || (param.type === 'array' ? '[1,2,3]' : param.type === 'object' ? '{"key":"value"}' : param.defaultValue?.toString())}
        multiline={param.type === 'array' || param.type === 'object'}
        rows={param.type === 'array' || param.type === 'object' ? 3 : 1}
        type={param.type === 'number' ? 'number' : 'text'}
      />
    );
  };

  return (
    <Card variant="outlined" sx={{ '&:hover': { boxShadow: 2 }, transition: 'box-shadow 0.2s' }}>
      <CardContent>
        <Stack spacing={2}>
          <Box>
            <Stack direction="row" spacing={1} alignItems="center" mb={1} flexWrap="wrap">
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
            ) : canOpenDirectly ? (
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
            ) : (
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
            )}
          </Stack>

          {hasParameters && (
            <Collapse in={expanded}>
              <Stack spacing={2} sx={{ pt: 1 }}>
                {parameters.length > 0 && (
                  <Box>
                    <Typography variant="subtitle2" fontWeight={600} mb={1} color="text.secondary">
                      Query Parameters
                    </Typography>
                    <Stack spacing={2}>
                      {parameters.map((param) =>
                        renderParameter(
                          param,
                          queryParams[param.name],
                          (val) => setQueryParams((prev) => ({ ...prev, [param.name]: val }))
                        )
                      )}
                    </Stack>
                  </Box>
                )}
                {bodyParameters.length > 0 && (
                  <Box>
                    <Typography variant="subtitle2" fontWeight={600} mb={1} color="text.secondary">
                      Request Body
                    </Typography>
                    <Stack spacing={2}>
                      {bodyParameters.map((param) =>
                        renderParameter(
                          param,
                          bodyParams[param.name],
                          (val) => setBodyParams((prev) => ({ ...prev, [param.name]: val }))
                        )
                      )}
                    </Stack>
                  </Box>
                )}
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
                maxHeight: 300,
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