import { Box, chakra, Flex, Heading } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import Rating from "../../components/Rating";
import api from "../../services/api";

export default function MelhoresCidades() {
  const [dados, adicionarDados] = useState([])

  useEffect(()  => {
    const fetch = async () => {
      const res = await api.get('cidades/melhores');
      adicionarDados(res.data)
    }

    fetch().catch(console.error);
  }, []);
  
  return (
    <Box mt={8}>
      <Heading as={'h3'} fontSize={'22'}>Melhores Cidades</Heading>
      <chakra.p mb={3} color={'gray.600'}>👛 Conheça as melhores cidades que tem as melhores avaliações</chakra.p>
      {dados?.dados?.map(i => (
        <Flex key={i?.id} maxW="md" mx="auto" my={4}>
          <Box>
            <chakra.h3 fontSize="1xl" fontWeight="semibold" color={"gray.800"}>{i?.nome}</chakra.h3>
            <Rating rating={i?.media_estrelas} numReviews={i?.qtd_avaliacoes} />
          </Box>
        </Flex>
      ))}
    </Box>
  );
}