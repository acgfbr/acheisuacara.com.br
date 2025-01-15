import { MantineProvider, Container, Title, Text, Box } from '@mantine/core';
import { Notifications } from '@mantine/notifications';
import URLShortenerForm from './components/URLShortenerForm';
import '@mantine/core/styles.css';
import '@mantine/notifications/styles.css';

function App() {
  return (
    <MantineProvider>
      <Notifications />
      <Box 
        style={{ 
          minHeight: '100vh',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          background: 'linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%)',
          padding: '1rem'
        }}
      >
        <Box 
          style={{
            width: '100%',
            maxWidth: '500px',
            backgroundColor: 'white',
            padding: '2rem',
            borderRadius: '1rem',
            boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)'
          }}
        >
          <Title order={1} ta="center" mb="lg" style={{ color: '#1a73e8' }}>
            Achei Sua Cara
          </Title>
          <Text c="dimmed" ta="center" mb="xl" size="lg">
            Shorten your marketplace URLs quickly and easily
          </Text>
          <URLShortenerForm />
        </Box>
      </Box>
    </MantineProvider>
  );
}

export default App;
