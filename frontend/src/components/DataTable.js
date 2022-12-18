import React from 'react'
import { Box, Button, Stack, Skeleton } from '@mui/material'
import { DataGrid } from '@mui/x-data-grid'
import { useState, useEffect } from 'react'


const columns = [
	{field: 'id', headerName: 'ID'},
	{field: 'title', headerName: 'Title', width:300},
	{field: 'body', headerName: 'Body', width:300}
]


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

	useEffect(() => {
		fetch("https://jsonplaceholder.typicode.com/posts")
		.then((data) => data.json())
		//.then((data) => console.log(data))
		.then((data) => {
			setTableData(data)
			setLoading(false)
		})
	});

	return (
		<div style={{height:650, width: '100%'}}>
		<DataGrid 
			rows={tableData}
			columns={columns}
			pageSize={10}
			rowsPerPageOptions={[10]}
			checkboxSelection
			components={{LoadingOverlay: LoadingSkeleton}}
			loading={false}
			onSelectionModelChange={(ids) => {
				const selectedIDs = new Set(ids);
				const selectedRows = tableData.rows.filter((row) =>
					selectedIDs.has(row.id),
				);
			setSelectedRows(selectedRows);
			}}
		{...tableData}
		/>
		<pre style={{ fontSize:10 }}>
		{JSON.stringify(selectedRows, null, 4)}
		</pre>

		<Stack spacing={2} direction="row" style={{ marginTop: '10px'}}>
	  	<Button variant="contained">Save</Button>
	  	<Button variant="contained">Process List</Button>
		</Stack>
		</div>
	);
}

export default DataTable
