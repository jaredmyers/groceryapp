import React from 'react';
import { useState, useEffect } from 'react'
import { Typography, AppBar, Toolbar, Container, CssBaseline, Box} from '@mui/material';
import { Card, CardActions, CardContent, CardMedia, Grid } from '@mui/material';
import Header from './components/Header';
import { DataGrid } from '@mui/x-data-grid'
import DataTable from './components/DataTable'



function App() {

  return (
	  <>
	  <CssBaseline />
	  <Header />

	  <Container maxWidth="sm" style={{ marginTop: '100px' }}>
	     <Typography variant="h3" align="center" color="textPrimary" gutterBottom>Create Grocery List</Typography>
	  <DataTable />
	  </Container>

	  </>
  );
}

export default App;
