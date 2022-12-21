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

	  <main> 
	   <div>
	    <Container maxWidth="md" style={{ marginTop: '100px'}}>
	     <Typography variant="h5" align="center" color="textSecondary" paragraph>
	  This is a test paragraph. This is for testing the test with all the more testing just to see how this will look.</Typography>

	     <div>
	  <Container maxWidth="md">
	      <Grid container spacing={4}>
	       <Grid item>
	  	<Card sx={{ minWidth:275 }}>
	  	<CardContent>
	  	  <Typography>testing</Typography>
	   	</CardContent>
	        </Card>
	       </Grid>
	      </Grid>
	      <Grid container spacing={4}>
	       <Grid item>
	  	<Card>
	  	<CardContent>
	  	  <Typography>testing</Typography>
	   	</CardContent>
	        </Card>
	       </Grid>
	      </Grid>
	  </Container>
	     </div>

	    </Container>
	   </div>
	  </main>

	  </>
  );
}

export default App;
