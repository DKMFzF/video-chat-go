import { createBrowserRouter } from "react-router-dom";

import MainPage from '../pages/Main';

const router = createBrowserRouter([
	{
		path: "/",
		element: <MainPage />,
	},
]);

export default router;
