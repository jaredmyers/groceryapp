import React from 'react'
import { Box, Button, Stack, Skeleton } from '@mui/material'
import { DataGrid } from '@mui/x-data-grid'
import { useState, useEffect, useRef } from 'react'
import axios from 'axios';


const columns = [
	{field: 'id', headerName: 'ID'},
	{field: 'title', headerName: 'Title', width:300},
	{field: 'body', headerName: 'Body', width:300}
]

const testcolumns = [
	{field: 'id', headerName: 'ID'},
	{field: 'item', headerName: 'Item', width:200}
]

var selection;


function useGetAllData(setTableData){
	useEffect(() => {
		var url = "http://localhost:8000/test"
		axios.get(url, {
			responseType: 'json'
		}).then(response => {
			if(response.status == 200){
				setTableData(response.data)
				console.log(response.data)
			}
		})
	}, [])
}


function processList(selection){
	
	var selectionJSON = JSON.stringify(selection)
	var url = "http://localhost:8000/add"
	axios.post(url, {
		"selections": selection
	}).then(response => {
		if(response.status == 200){
			alert("list processed")
		}
	})

}

function manageChecks(itm){
	selection = itm;
	console.log(selection)
}

function saveList(){

}


const LoadingSkeleton = () => (
	<Box
	  sx={{
		height: "max-content"
	  }}
	>
	{[...Array(10)].map((_) => (
		<Skeleton variant="rectangular" height={45} sx= {{my:1, mx:0}}/>
	))}
	</Box>
);

const DataTable = () => {
	const [tableData, setTableData] = useState([]);
	const [loading, setLoading] = useState(true);
	const [selectedRows, setSelectedRows] = useState([]);
	
	useGetAllData(setTableData)	
	/*
	useEffect(() => {
		//fetch("https://jsonplaceholder.typicode.com/posts")
		fetch("http://localhost:8000/test")
		.then((data) => data.json())
		//.then((data) => console.log(data))
		.then((data) => {
			console.log(data)
			setTableData(data)
			setLoading(false)
		})
	}, []);
	*/
	return (
		<div style={{height:650, width: '100%'}}>
		<DataGrid 
			rows={tableData}
			columns={testcolumns}
			pageSize={10}
			rowsPerPageOptions={[10]}
			checkboxSelection
			components={{LoadingOverlay: LoadingSkeleton}}
			loading={false}
			onSelectionModelChange={itm => manageChecks(itm)}
		/>
		<Stack spacing={2} direction="row" style={{ marginTop: '10px'}}>
	  	<Button variant="contained">Save</Button>
	  	<Button variant="contained"
			onClick={() => {processList(selection);}}>Process List</Button>
		</Stack>
		</div>
	);
}

export default DataTable
