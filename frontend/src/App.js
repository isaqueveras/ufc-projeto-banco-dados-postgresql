import { Routes, Route } from "react-router-dom"

import TelaPrincipal from "./pages/Inicio/Inicio";
import DashboadInicio from "./pages/Dashboard/dashboardInicio";

import TelaCategorias from "./pages/Dashboard/Categorias";
import EditarCategorias from "./pages/Dashboard/Categorias/editar";
import CadastrarCategorias from "./pages/Dashboard/Categorias/cadastrar";

import TelaCidades from "./pages/Dashboard/Cidades";
import EditarCidades from "./pages/Dashboard/Cidades/editar";
import CadastrarCidade from "./pages/Dashboard/Cidades/cadastrar";

import TelaAvaliacoes from "./pages/Dashboard/Avaliacoes";
import EditarAvaliacoes from "./pages/Dashboard/Avaliacoes/editar";
import CadastrarAvaliacoes from "./pages/Dashboard/Avaliacoes/cadastrar";

import TelaEmpresas from "./pages/Dashboard/Empresas";
import EditarEmpresa from "./pages/Dashboard/Empresas/editar";
import CadastrarEmpresa from "./pages/Dashboard/Empresas/cadastrar";

function App() {
  return (
    <Routes>
      <Route index path="/" element={ <TelaPrincipal /> } />
      <Route path="dashboard" element={ <DashboadInicio /> } />

      <Route path="dashboard/empresas" element={ <TelaEmpresas /> } />
      <Route path="dashboard/empresas/editar/:id" element={ <EditarEmpresa /> } />
      <Route path="dashboard/empresas/cadastrar" element={ <CadastrarEmpresa /> } />
      
      <Route path="dashboard/cidades" element={ <TelaCidades /> } />
      <Route path="dashboard/cidades/editar/:id" element={ <EditarCidades /> } />
      <Route path="dashboard/cidades/cadastrar" element={ <CadastrarCidade /> } />
      
      <Route path="dashboard/categorias" element={ <TelaCategorias /> } />
      <Route path="dashboard/categorias/editar/:id" element={ <EditarCategorias /> } />
      <Route path="dashboard/categorias/cadastrar" element={ <CadastrarCategorias /> } />
      
      <Route path="dashboard/avaliacoes" element={ <TelaAvaliacoes /> } />
      <Route path="dashboard/avaliacoes/editar/:id" element={ <EditarAvaliacoes /> } />
      <Route path="dashboard/avaliacoes/cadastrar" element={ <CadastrarAvaliacoes /> } />
      
    </Routes>
  );
}

export default App;
