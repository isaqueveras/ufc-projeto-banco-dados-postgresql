import { Box, chakra, Heading, Tag } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import Rating from "../../components/Rating";
import api from "../../services/api";

export default function UltimasAvaliacoes() {
  const [dados, adicionarDados] = useState([])

  useEffect(()  => {
    const fetch = async () => {
      const res = await api.get('avaliacoes/ultimas');
      adicionarDados(res.data)
    }

    fetch().catch(console.error);
  }, []);
  
  return (
    <Box>
      {dados?.dados?.map(i => (
         <Box 
            backgroundColor={'whiteAlpha.600'}
            boxShadow={'sm'}
            my='3'
            pb='3'
            mt={8}
            borderBottom='2px solid rgba(0,0,0,0.08)'
            key={i?.id}
          >
            <Heading fontSize={23}>{i?.titulo}</Heading>
            <chakra.p color={'gray.700'} my='1'><i>{i?.nome_pessoa}</i></chakra.p>
            <chakra.p color={'gray.700'} mb='2'>{i?.descricao}</chakra.p>
            <Rating rating={i?.qtd_estrela} numReviews={0} />
            <Tag size={'md'} variant='subtle' colorScheme='green' mt='3' mr='2'>{i?.categoria}</Tag>
          </Box>
      ))}
    </Box>
  );
}