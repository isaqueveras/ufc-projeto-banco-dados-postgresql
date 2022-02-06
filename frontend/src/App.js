import { Routes, Route } from "react-router-dom"

import TelaPrincipal from "./pages/Inicio";
import TelaEmpresas from "./pages/Empresas/Inicio";
import TelaCategorias from "./pages/Categorias/Inicio";
import TelaCidades from "./pages/Cidades/Inicio";
import TelaAvaliacoes from "./pages/Avaliacoes/Inicio";

function App() {
  return (
    <Routes>
      <Route index path="/" element={ <TelaPrincipal /> } />
      <Route path="cidades" element={ <TelaCidades /> } />
      <Route path="empresas" element={ <TelaEmpresas /> } />
      <Route path="categorias" element={ <TelaCategorias /> } />
      <Route path="avaliacoes" element={ <TelaAvaliacoes /> } />
    </Routes>
  );
}

export default App;
