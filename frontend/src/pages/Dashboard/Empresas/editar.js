import { Alert, AlertIcon, Button, chakra, CloseButton, Container, Grid, GridItem, Heading, Input } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";

import MenuDashboard from "../../../components/MenuDashboard";
import api from "../../../services/api";

export default function EditarEmpresa() {
	const { id } = useParams()

	const [nome, adicionarNome] = useState()
	const [telefone, adicionarTelefone] = useState()
	const [cpf, adicionarCPF] = useState()
	const [cnpj, adicionarCNPJ] = useState()
	const [cidade_id, adicionarCidadeID] = useState(0)
	const [categoria_id, adicionarCategoriaID] = useState(0)
	const [dadosRec, adicionarDadosRec] = useState({})

  const [mensagem, adicionarMensagem] = useState({})

  useEffect(()  => {
		const fetchData = async () => {
			// Busca uma lista de empresa utilizando a API feito em Golang
			const res = await api.get('empresa/'+id);

			adicionarNome(res?.data?.nome)
			adicionarTelefone(res?.data?.telefone)
			adicionarCPF(res?.data?.cpf)
			adicionarCNPJ(res?.data?.cnpj)
			adicionarCidadeID(res?.data?.cidade_id)
			adicionarCategoriaID(res?.data?.categoria_id)

			adicionarDadosRec(res.data)
		}
		
    fetchData().catch(console.error);
  }, [id]);

	var infoEditado = {
		nome: nome !== '' ? nome : dadosRec?.nome,
		telefone: telefone !== '' ? telefone : dadosRec?.telefone,
		cpf: cpf !== '' ? cpf : dadosRec?.cpf,
		cnpj: cnpj !== '' ? cnpj : dadosRec?.cnpj,
		cidade_id: cidade_id !== '' ? parseInt(cidade_id) : parseInt(dadosRec?.cidade_id),
		categoria_id: categoria_id !== '' ? parseInt(categoria_id) : parseInt(dadosRec?.categoria_id),
	}
	
  function editarEmpresa() {
    api.put(`empresa/${id}`, infoEditado).then(res => {
      adicionarMensagem({
        tipo: 'success', 
        mensagem: `Empresa com o id (${id}) foi editado com sucesso.`,
        estaAtivo: true,
      });
    }).catch(error => {
      adicionarMensagem({
        tipo: 'error',
        mensagem: `Erro ao editar a empresa com o id (${id}).`,
        estaAtivo: true
      });
    });
  }

	return (
		<>
			<MenuDashboard />
			<Container maxW='8xl' my={8}>
				{mensagem?.estaAtivo === true && (
          <Alert status={mensagem?.tipo} variant='top-accent' my={5}>
            <AlertIcon />{mensagem?.mensagem}<CloseButton onClick={() => adicionarMensagem({estaAtivo: false})} position='absolute' right='8px'/>
          </Alert>
        )}
				
				<Heading mb={5}>Editar dados da empresa</Heading>

				<Grid templateColumns='repeat(3, 1fr)' gap={6}>
					<GridItem w='100%'>
						<chakra.b>Nome da empresa</chakra.b>
						<Input variant='outline' placeholder='Nome' value={nome} onChange={e => adicionarNome(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						<chakra.b>Telefone</chakra.b>
						<Input variant='outline' placeholder='Telefone' value={telefone} onChange={e => adicionarTelefone(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						<chakra.b>CPF</chakra.b>
						<Input variant='outline' placeholder='CPF' value={cpf} onChange={e => adicionarCPF(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						<chakra.b>CNPJ</chakra.b>
						<Input variant='outline' placeholder='CNPJ' value={cnpj} onChange={e => adicionarCNPJ(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						<chakra.b>ID Cidade</chakra.b>
						<Input variant='outline' placeholder='ID Cidade' value={cidade_id} onChange={e => adicionarCidadeID(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						<chakra.b>ID Categoria</chakra.b>
						<Input variant='outline' placeholder='ID Categoria' value={categoria_id} onChange={e => adicionarCategoriaID(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						<Button colorScheme='blue' size='md' onClick={() => editarEmpresa()}>Editar dados</Button>{' '}
						<Link to="../dashboard/empresas"><Button colorScheme='gray' size='md'>Cancelar</Button></Link>
					</GridItem>
				</Grid>
			</Container>
		</>
	)
}