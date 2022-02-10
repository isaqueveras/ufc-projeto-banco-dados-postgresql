import { Box, chakra, Container, Grid, GridItem, useColorModeValue } from "@chakra-ui/react";
import MenuDashboard from "../../components/MenuDashboard";

export default function DashboadInicio() {
	return (
		<>
			<MenuDashboard />
      <Container maxW='8xl'>
				<Grid templateColumns='repeat(1, 1fr)' gap={0}>
					<GridItem>
						<Box py={12}>
							<chakra.p mb={2} fontSize="xs" fontWeight="semibold" letterSpacing="wide" color="gray.400" textTransform="uppercase">Para administradores</chakra.p>
							<chakra.h1 fontSize={{ base: "3xl", md: "4xl" }} fontWeight="bold" lineHeight="shorter" color={useColorModeValue("gray.900", "white")}>🎉 Gerencie o sistema por aqui :)</chakra.h1>
							<chakra.p mb={2} color="gray.500" fontSize={{ md: "lg" }}>
								Somos um centro de avaliações onde os consumidores podem encontrar todo tipo de avalições. <br /> Avaliações de comerciantes, prestadores de serviços, empresas grandes, produtos e etc.
							</chakra.p>
						</Box>
					</GridItem>
				</Grid>
			</Container>
		</>
	)
}