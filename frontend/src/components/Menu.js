import { Box, Flex, HStack, Link as LinkItem, Container, Heading } from '@chakra-ui/react';
import { Link } from "react-router-dom"

export default function Menu() {
  return (
    <Box bg={'gray.800'} px={4}>
      <Container maxW='8xl'>
        <Flex h={16} alignItems={'center'} justifyContent={'space-between'}>
          <HStack spacing={8} alignItems={'center'}>
            <Box><Heading color={'white'} fontSize={22}>Centro de Avaliações</Heading></Box>
            <HStack as={'nav'} spacing={4} display={{ base: 'none', md: 'flex' }}>
              <LinkItem
                px={2}
                py={1}
                rounded={'md'}
                _hover={{ textDecoration: 'none', bg: 'messenger.800', textDecor: 'none'}}
                color={'gray.200'}>
                <Link to="dashboard">✏️ Dashboard</Link>
              </LinkItem>
            </HStack>
          </HStack>
        </Flex>
      </Container>
    </Box>
  );
}
