import { Box, Flex, HStack, Link as LinkItem, Container, Heading } from '@chakra-ui/react';
import { Link } from "react-router-dom"

export default function MenuDashboard() {
  return (
    <Box bg={'gray.800'} px={4}>
      <Container maxW='8xl'>
        <Flex h={16} alignItems={'center'} justifyContent={'space-between'}>
          <HStack spacing={8} alignItems={'center'}>
            <Box><Heading color={'white'} fontSize={22}><Link to="/dashboard">Dashbaord</Link></Heading></Box>
            <HStack as={'nav'} spacing={4} display={{ base: 'none', md: 'flex' }}>
              <LinkItem
                px={2}
                py={1}
                rounded={'md'}
                _hover={{ textDecoration: 'none', bg: 'messenger.800', textDecor: 'none'}}
                color={'gray.200'}>
                <Link to="/">Tela inicial</Link>
              </LinkItem>
              <LinkItem
                px={2}
                py={1}
                rounded={'md'}
                _hover={{ textDecoration: 'none', bg: 'messenger.800', textDecor: 'none'}}
                color={'gray.200'}>
                <Link to="cidades">Cidades</Link>
              </LinkItem>
              <LinkItem
                px={2}
                py={1}
                rounded={'md'}
                _hover={{ textDecoration: 'none', bg: 'messenger.800', textDecor: 'none'}}
                color={'gray.200'}>
                <Link to="empresas">Empresas</Link>
              </LinkItem>
              <LinkItem
                px={2}
                py={1}
                rounded={'md'}
                _hover={{ textDecoration: 'none', bg: 'messenger.800', textDecor: 'none'}}
                color={'gray.200'}>
                <Link to="categorias">Categorias</Link>
              </LinkItem>
              <LinkItem
                px={2}
                py={1}
                rounded={'md'}
                _hover={{ textDecoration: 'none', bg: 'messenger.800', textDecor: 'none'}}
                color={'gray.200'}>
                <Link to="avaliacoes">Avaliações</Link>
              </LinkItem>
            </HStack>
          </HStack>
        </Flex>
      </Container>
    </Box>
  );
}
