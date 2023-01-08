import React from "react";
import { Box, Button, TextField, Typography, CssBaseline } from "@mui/material";

const Auth = () => {
		return (
				<>
				<CssBaseline />
				<Typography variant="h3" align="center" color="textPrimary">Login</Typography>
				 
				 <form>
				  <Box 
				    display="flex" 
				    flexDirection={"column"}
				    maxWidth={400}
				    alignItems={"center"}
				    justifyContent={"center"}
				    margin="auto"
				    marginTop={1}
				    padding={3}
				borderRadius={5}>

				   <TextField placeholder="Username"/>
				   <TextField placeholder="Password"/>
				   <TextField />
				   <Button>Login</Button>
				</Box>
				 </form>
				</>
		);
};

export default Auth;

