import { Link } from "react-router-dom"

export default function Menu() {
  return (
    <div>
      <Link to="">Pagina Principal</Link>{' '}
      <Link to="empresas">Pagina Empresas</Link>{' '}
      <Link to="cidades">Pagina Cidades</Link>{' '}
      <Link to="categorias">Pagina Categorias</Link>{' '}
      <Link to="avaliacoes">Pagina Avaliacoes</Link>{' '}
    </div>
  )
}