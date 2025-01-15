import { useState } from 'react';
import { TextInput, Button, Paper, CopyButton, Text, Stack, Box } from '@mantine/core';
import { notifications } from '@mantine/notifications';
import { IconLink, IconCopy, IconCheck } from '@tabler/icons-react';
import axios from 'axios';

const API_URL = 'http://localhost:8080';

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
      const response = await axios.post(`${API_URL}/api/shorten`, { url });
      setShortenedURL(response.data);
      setUrl('');
      notifications.show({
        title: 'Success',
        message: 'URL shortened successfully!',
        color: 'green',
      });
    } catch (error: any) {
      notifications.show({
        title: 'Error',
        message: error.response?.data?.error || 'Failed to shorten URL',
        color: 'red',
      });
    } finally {
      setLoading(false);
    }
  };

  const shortUrl = shortenedURL ? `${API_URL}/${shortenedURL.short_code}` : '';

  return (
    <Stack gap="lg">
      <form onSubmit={handleSubmit}>
        <TextInput
          required
          size="lg"
          label="Enter marketplace URL"
          placeholder="https://www.amazon.com/product/123"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          mb="md"
          leftSection={<IconLink size={20} style={{ color: '#1a73e8' }} />}
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
          Shorten URL
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
            Shortened URL:
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
                leftSection={copied ? <IconCheck size={20} /> : <IconCopy size={20} />}
                variant={copied ? 'light' : 'filled'}
                style={{
                  backgroundColor: copied ? undefined : '#1a73e8',
                }}
              >
                {copied ? 'Copied!' : 'Copy URL'}
              </Button>
            )}
          </CopyButton>
        </Paper>
      )}
    </Stack>
  );
} 