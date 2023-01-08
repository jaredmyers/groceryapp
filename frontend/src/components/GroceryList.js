import React from "react";
import { Container, Box, Button, TextField, Typography } from "@mui/material";
import DataTable from './DataTable'
import Header from "./Header"

const GroceryList = () => {
		return (
				<>
				<Header />
				<Container maxWidth="sm" style={{ marginTop: '100px' }}>
				<Typography variant="h3" align="center" color="textPrimary" gutterBottom>Create Grocery List</Typography>
				<DataTable />
				</Container>
				</>
		);
};

export default GroceryList;
