import { useState } from 'react';
import { TextInput, Button, Paper, CopyButton, Text, Stack, Box } from '@mantine/core';
import { notifications } from '@mantine/notifications';
import { FiLink, FiCopy, FiCheck } from 'react-icons/fi';
import axios from 'axios';
import { config } from '../config';

// Pre-configure axios
const api = axios.create({
  baseURL: config.apiUrl,
  timeout: 5000
});

interface ShortenedURL {
  url: string;
  short_code: string;
}

export default function URLShortenerForm() {
  const [url, setUrl] = useState('');
  const [shortenedURL, setShortenedURL] = useState<ShortenedURL | null>(null);
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      const response = await api.post('/api/shorten', { url });
      setShortenedURL(response.data);
      setUrl('');
      notifications.show({
        title: 'Success',
        message: 'URL encurtada com sucesso!',
        color: 'green',
      });
    } catch (error: any) {
      notifications.show({
        title: 'Error',
        message: error.response?.data?.error || 'Falhou pra encurtar a URL meu capitão',
        color: 'red',
      });
    } finally {
      setLoading(false);
    }
  };

  const shortUrl = shortenedURL ? `${config.apiUrl}/${shortenedURL.short_code}` : '';

  return (
    <Stack gap="lg">
      <form onSubmit={handleSubmit}>
        <TextInput
          required
          size="lg"
          label="Joga aqui o link"
          placeholder="https://www.amazon.com/product/123"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          mb="md"
          leftSection={<FiLink size={20} style={{ color: '#1a73e8' }} />}
          styles={{
            input: {
              '&:focus': {
                borderColor: '#1a73e8',
              },
            },
          }}
        />
        <Button
          type="submit"
          loading={loading}
          fullWidth
          size="lg"
          style={{
            backgroundColor: '#1a73e8',
            '&:hover': {
              backgroundColor: '#1557b0',
            },
          }}
        >
          Encurtar
        </Button>
      </form>

      {shortenedURL && (
        <Paper
          withBorder
          p="lg"
          radius="md"
          style={{
            backgroundColor: '#f8f9fa',
            borderColor: '#e9ecef',
          }}
        >
          <Text size="sm" fw={500} mb="xs" c="dimmed">
            URL encurtada:
          </Text>
          <Box
            style={{
              backgroundColor: 'white',
              padding: '0.75rem',
              borderRadius: '0.5rem',
              border: '1px solid #e9ecef',
              marginBottom: '1rem',
              wordBreak: 'break-all',
            }}
          >
            <Text size="md" style={{ color: '#1a73e8' }}>
              {shortUrl}
            </Text>
          </Box>
          <CopyButton value={shortUrl}>
            {({ copied, copy }) => (
              <Button
                color={copied ? 'teal' : 'blue'}
                onClick={copy}
                leftSection={copied ? <FiCheck size={20} /> : <FiCopy size={20} />}
                variant={copied ? 'light' : 'filled'}
                style={{
                  backgroundColor: copied ? undefined : '#1a73e8',
                }}
              >
                {copied ? 'Copiou!' : 'Copiar URL'}
              </Button>
            )}
          </CopyButton>
        </Paper>
      )}
    </Stack>
  );
} 