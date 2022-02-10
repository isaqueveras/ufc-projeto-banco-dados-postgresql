import { Alert, AlertIcon, Button, chakra, CloseButton, Container, Grid, GridItem, Heading, Input } from "@chakra-ui/react";
import { useState } from "react";
import { Link } from "react-router-dom";

import MenuDashboard from "../../../components/MenuDashboard";
import api from "../../../services/api";

export default function CadastrarAvaliacao() {
  const [mensagem, adicionarMensagem] = useState({})
	
	const [titulo, adicionarTitulo] = useState()
	const [descricao, adicionarDescricao] = useState()
	const [qtd_estrela, adicionarQtdEstrelas] = useState()
	const [nome_pessoa, adicionarNomePessoa] = useState()
	const [empresa_id, adicionarEmpresaID] = useState()

	var dadosCadastrar = {
		titulo: titulo,
		descricao: descricao,
		qtd_estrela: parseFloat(qtd_estrela),
		nome_pessoa: nome_pessoa,
		empresa_id: parseInt(empresa_id)
	}
	
  function cadastrarAvaliacao() {
    api.post(`avaliacoes`, dadosCadastrar).then(res => {
      adicionarMensagem({
        tipo: 'success', 
        mensagem: `Avaliação cadastrada com sucesso.`,
        estaAtivo: true,
      });

			adicionarTitulo(''); adicionarDescricao(''); adicionarQtdEstrelas(); adicionarNomePessoa();
    }).catch(error => {
      adicionarMensagem({
        tipo: 'error',
        mensagem: `Erro ao cadastrar a avaliação.`,
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
				
				<Heading mb={5}>Cadastrar uma nova avaliação</Heading>

				<Grid templateColumns='repeat(2, 1fr)' gap={6}>
					<GridItem w='100%'>
						<chakra.b>Titulo</chakra.b>
						<Input variant='outline' placeholder='Titulo' value={titulo} onChange={e => adicionarTitulo(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						<chakra.b>Descrição</chakra.b>
						<Input variant='outline' placeholder='Descrição' value={descricao} onChange={e => adicionarDescricao(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						<chakra.b>Quantidade Estrelas</chakra.b>
						<Input variant='outline' placeholder='Quantidade Estrelas' value={qtd_estrela} onChange={e => adicionarQtdEstrelas(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						<chakra.b>Nome da pessoa</chakra.b>
						<Input variant='outline' placeholder='Nome da pessoa' value={nome_pessoa} onChange={e => adicionarNomePessoa(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						<chakra.b>Empresa ID</chakra.b>
						<Input variant='outline' placeholder='ID da empresa' value={empresa_id} onChange={e => adicionarEmpresaID(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						{titulo?.length > 5 && descricao !== '' ? (
							<Button colorScheme='blue' size='md' onClick={() => cadastrarAvaliacao()}>Cadastrar Avaliação</Button>
						) : (
							<Button colorScheme='gray' size='md'>Cadastrar Avaliação</Button>
						)}
						{' '}<Link to="../dashboard/cidades"><Button colorScheme='gray' size='md'>Cancelar</Button></Link>
					</GridItem>
				</Grid>
			</Container>
		</>
	)
}