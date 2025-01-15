import { MantineProvider, Title, Text, Box } from '@mantine/core';
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
          width: '100vw',
          height: '100vh',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          background: 'linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%)',
          margin: 0,
          padding: 0,
          position: 'fixed',
          top: 0,
          left: 0
        }}
      >
        <Box 
          style={{
            width: '90%',
            maxWidth: '500px',
            backgroundColor: 'white',
            padding: '2rem',
            borderRadius: '1rem',
            boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
            margin: 'auto'
          }}
        >
          <Title order={1} ta="center" mb="lg" style={{ color: '#1a73e8' }}>
            Achei Sua Cara
          </Title>
          <Text c="dimmed" ta="center" mb="xl" size="lg">
            Encurte suas URLs de marketplace rapidamente!
          </Text>
          <URLShortenerForm />
        </Box>
      </Box>
    </MantineProvider>
  );
}

export default App;
