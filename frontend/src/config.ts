// Default to production URL if environment variable is not set
const API_URL = import.meta.env.VITE_API_URL || 'https://api.acheisuacara.com.br';
const FRONTEND_URL = import.meta.env.VITE_FRONTEND_URL || 'https://acheisuacara.com.br';
export const config = {
  apiUrl: API_URL,
  frontendUrl: FRONTEND_URL
}; 