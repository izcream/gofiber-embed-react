import { RouterProvider, createBrowserRouter } from 'react-router-dom'
import AppLayout from './components/Layout'

const Home = () => <h1 className="text-4xl text-center">Home Page</h1>
const About = () => <h1 className="text-4xl text-center">About Page</h1>
const routes = createBrowserRouter([
  {
    path: '/',
    element: <AppLayout />,
    children: [
      {
        path: '',
        element: <Home />
      },
      {
        path: '/about',
        element: <About />
      }
    ]
  }
])

function App() {
  return <RouterProvider router={routes} />
}

export default App
