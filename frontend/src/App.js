import { Routes, Route } from "react-router-dom"

import TelaPrincipal from "./pages/Inicio/Inicio";
import DashboadInicio from "./pages/Dashboard/dashboardInicio";
import TelaEmpresas from "./pages/Dashboard/Empresas";
import TelaCategorias from "./pages/Dashboard/Categorias/Inicio";
import TelaCidades from "./pages/Dashboard/Cidades/Inicio";
import TelaAvaliacoes from "./pages/Dashboard/Avaliacoes/Inicio";

function App() {
  return (
    <Routes>
      <Route index path="/" element={ <TelaPrincipal /> } />
      <Route path="dashboard" element={ <DashboadInicio /> } />
      <Route path="/dashboard/empresas" element={ <TelaEmpresas /> } />
      <Route path="/dashboard/cidades" element={ <TelaCidades /> } />
      <Route path="/dashboard/categorias" element={ <TelaCategorias /> } />
      <Route path="/dashboard/avaliacoes" element={ <TelaAvaliacoes /> } />
    </Routes>
  );
}

export default App;
