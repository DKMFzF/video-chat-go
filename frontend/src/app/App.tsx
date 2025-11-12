import { StrictMode } from "react";
import { HelmetProvider } from "react-helmet-async";
import { RouterProvider } from "react-router-dom";

import router from "./routes";

export const App = (): React.JSX.Element => {
	return (
		<StrictMode>
			<HelmetProvider>
				<RouterProvider router={router} />
			</HelmetProvider>
		</StrictMode>
	);
};
