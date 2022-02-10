import { Alert, AlertIcon, Button, chakra, CloseButton, Container, Grid, GridItem, Heading, Input } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";

import MenuDashboard from "../../../components/MenuDashboard";
import api from "../../../services/api";

export default function EditarAvaliacao() {
	const { id } = useParams()

	const [titulo, adicionarTitulo] = useState()
	const [descricao, adicionarDescricao] = useState()
	const [qtd_estrela, adicionarQtdEstrelas] = useState()
	const [nome_pessoa, adicionarNomePessoa] = useState()

	const [dadosRec, adicionarDadosRec] = useState({})
  const [mensagem, adicionarMensagem] = useState({})

  useEffect(()  => {
		const fetchData = async () => {
			// Busca os dados de uma avaliação utilizando a API feito em Golang
			const res = await api.get('avaliacao/'+id);

			adicionarTitulo(res?.data.nome)
			adicionarDescricao(res?.data?.descricao)
			adicionarQtdEstrelas(res?.data?.qtd_estrela)
			adicionarNomePessoa(res?.data?.nome_pessoa)
			adicionarDadosRec(res.data)
		}
		
    fetchData().catch(console.error);
  }, [id]);

	var infoEditado = {
		titulo: titulo !== '' ? titulo : dadosRec?.titulo,
		descricao: descricao !== '' ? descricao : dadosRec?.descricao,
		qtd_estrela: qtd_estrela !== '' ? qtd_estrela : dadosRec?.qtd_estrela,
		nome_pessoa: nome_pessoa !== '' ? nome_pessoa : dadosRec?.nome_pessoa,
	}
	
  function editarAvaliacao() {
    api.put(`avaliacao/${id}`, infoEditado).then(res => {
      adicionarMensagem({
        tipo: 'success', 
        mensagem: `Avaliação com o id (${id}) foi editado com sucesso.`,
        estaAtivo: true,
      });
    }).catch(error => {
      adicionarMensagem({
        tipo: 'error',
        mensagem: `Erro ao editar a avaliação com o id (${id}).`,
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
				
				<Heading mb={5}>Editar dados da avaliação</Heading>
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
						<chakra.b>Cidade ID</chakra.b>
						<Input variant='outline' placeholder='ID da cidade' value={nome_pessoa} onChange={e => adicionarNomePessoa(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						<Button colorScheme='blue' size='md' onClick={() => editarAvaliacao()}>Editar dados</Button>{' '}
						<Link to="../dashboard/categorias"><Button colorScheme='gray' size='md'>Cancelar</Button></Link>
					</GridItem>
				</Grid>
			</Container>
		</>
	)
}