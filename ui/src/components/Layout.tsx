import { Link, Outlet } from 'react-router-dom'

function AppLayout() {
  return (
    <>
      <nav className="container mx-auto h-[70px] flex items-center justify-between">
        <Link to="/">GoFiber+React</Link>
        <div className="space-x-6">
          <Link to="/">Home</Link>
          <Link to="/about">About</Link>
        </div>
      </nav>
      <div className="container mx-auto py-20">
        <Outlet />
      </div>
    </>
  )
}

export default AppLayout
