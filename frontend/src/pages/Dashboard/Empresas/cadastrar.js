import { Alert, AlertIcon, Button, chakra, CloseButton, Container, Grid, GridItem, Heading, Input } from "@chakra-ui/react";
import { useState } from "react";
import { Link } from "react-router-dom";

import MenuDashboard from "../../../components/MenuDashboard";
import api from "../../../services/api";

export default function CadastrarEmpresa() {
  const [mensagem, adicionarMensagem] = useState({})
	
	const [nome, adicionarNome] = useState()
	const [telefone, adicionarTelefone] = useState()
	const [cpf, adicionarCPF] = useState()
	const [cnpj, adicionarCNPJ] = useState()
	const [cidade_id, adicionarCidadeID] = useState(0)
	const [categoria_id, adicionarCategoriaID] = useState(0)

	var dadosCadastrar = {
		nome: nome,
		telefone: telefone,
		cpf: cpf,
		cnpj: cnpj,
		cidade_id: parseInt(cidade_id),
		categoria_id: parseInt(categoria_id),
	}
	
  function cadastrarEmpresa() {
    api.post(`empresas`, dadosCadastrar).then(res => {
      adicionarMensagem({
        tipo: 'success', 
        mensagem: `Empresa cadastrada com sucesso.`,
        estaAtivo: true,
      });

			adicionarNome(''); adicionarTelefone(''); adicionarCPF('')
			adicionarCNPJ(''); adicionarCidadeID(0); adicionarCategoriaID(0)
    }).catch(error => {
      adicionarMensagem({
        tipo: 'error',
        mensagem: `Erro ao cadastrar a empresa.`,
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
				
				<Heading mb={5}>Cadastrar uma nova empresa</Heading>

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
						{nome?.length > 5 && telefone !== '' && cidade_id !== '' & categoria_id  !== '' ? (
							<Button colorScheme='blue' size='md' onClick={() => cadastrarEmpresa()}>Cadastrar Empresa</Button>
						) : (
							<Button colorScheme='gray' size='md'>Cadastrar Empresa</Button>
						)}
						{' '}<Link to="../dashboard/empresas"><Button colorScheme='gray' size='md'>Cancelar</Button></Link>
					</GridItem>
				</Grid>
			</Container>
		</>
	)
}