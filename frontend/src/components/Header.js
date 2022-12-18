import { Typography, AppBar, Toolbar, CssBaseline } from '@mui/material';
import { Stack, Button } from '@mui/material';

const Header = () => {

	return (
		<>
		<CssBaseline />
		<AppBar postion="relative">
		<Toolbar>
		<Stack spacing={2} direction="row">
		<Typography variant="h4">GroceryApp</Typography>
		<Button variant="text" color="secondary">Add Item</Button>
		<Button variant="text" color="secondary">Shopping List</Button>
		<Button variant="text" color="secondary">Show Path</Button>
		</Stack>

		</Toolbar>
		</AppBar>
		</>
	);
}

export default Header
