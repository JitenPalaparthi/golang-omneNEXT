import NavBar from './components/NavBar';
import HomePage from './pages/HomePage'
import RegistrationPage from './pages/RegistrationPage'
import LocalUsersPage from './pages/LocalUsersPage'
import { Route, Routes } from "react-router-dom";

export default function App() {
  return (
    <>
      <NavBar />
      <div className="mx-auto max-w-5xl px-4 py-6">
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/users" element={<LocalUsersPage />} />
          <Route path="/registration" element={<RegistrationPage />} />
          <Route path="*" element={<h2 className="p-6">404 — Not Found</h2>} />
        </Routes>
      </div>
    </>
  );
}
