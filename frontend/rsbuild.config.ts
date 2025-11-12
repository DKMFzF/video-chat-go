import { defineConfig } from '@rsbuild/core';
import { pluginSvgr } from '@rsbuild/plugin-svgr';
import { pluginReact } from '@rsbuild/plugin-react';
import {
  pluginCssMinimizer,
  CssMinimizerWebpackPlugin,
} from '@rsbuild/plugin-css-minimizer';

export default defineConfig({
	plugins: [
		pluginReact(),
		pluginCssMinimizer({
      pluginOptions: {
        minify: CssMinimizerWebpackPlugin.lightningCssMinify,
      },
    }),
		pluginSvgr(),
	],

	source: {
		define: {
			BASE_URI: JSON.stringify(process.env.BASE_URI),
			JWT: JSON.stringify(process.env.JWT),
		}
	},

	resolve: {
		alias: {
			'@': './src',
			'@app': './src/app',
			'@entities': './src/entities',
			'@features': './src/features',
			'@widgets': './src/widgets',
			'@ui': './src/shared/ui',
			'@store': './src/shared/lib/store',
			'@api': './src/shared/api/base/HttpClient.ts',
			'@shared/lib': './src/shared/lib',
			'@api/types': './src/shared/api',
			'@api/error/access': './src/shared/api/base/error/NoAccess.ts',
			'@model': './src/shared/model',
			'@bem': './src/shared/lib/bem/cn.ts',
			'@assets': './src/assets',
		}
	},

	output: {
		minify: true,
	},

	server: {
		host: '0.0.0.0'
	}

	// server: {
	// 	port: process.env.PORT ? Number(process.env.PORT) : 8081,
	// 	// proxy: {
	// 	// 	'/api': {
	// 	// 		target: 'http://localhost:8080',
	// 	// 		changeOrigin: true,
	// 	// 	},
	// 	// }
	// },
});
