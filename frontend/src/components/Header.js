import { Typography, AppBar, Toolbar, CssBaseline } from '@mui/material';
import { Stack, Button } from '@mui/material';
import { Link } from 'react-router-dom'

const Header = () => {

	return (
		<>
		<CssBaseline />
		<AppBar postion="relative">
		<Toolbar>
		<Stack spacing={2} direction="row">
		<Typography variant="h4">GroceryApp</Typography>
		<Button variant="text" color="secondary"><Link to="/dashboard">Dashboard</Link></Button>
		<Button variant="text" color="secondary"><Link to="/grocerylist">Grocery List</Link></Button>
		<Button variant="text" color="secondary"><Link to="/showpath">Show Path</Link></Button>
		<Button variant="text" color="secondary"><Link to="/newitem">New Item</Link></Button>
		</Stack>

		</Toolbar>
		</AppBar>
		</>
	);
}

export default Header
