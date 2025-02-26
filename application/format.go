package main

func formatTextToHTML(text string) string {
	var html = `
				<!DOCTYPE html>
				<html lang="en">
				<head>
					<meta charset="UTF-8">
					<meta name="viewport" content="width=device-width, initial-scale=1.0">
					<title>Environment Variable Card</title>
					<style>
						body {
							font-family: Arial, sans-serif;
							display: flex;
							justify-content: center;
							align-items: center;
							height: 100vh;
							margin: 0;
							background-color: #f4f4f9;
						}
						.card {
							background-color: #ffffff;
							padding: 20px;
							border-radius: 8px;
							box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
							max-width: 300px;
							text-align: center;
						}
						.card h3 {
							color: #333;
						}
					</style>
				</head>
				<body>
					<div class="card">
						<h1>App works...</h1>
						<h3>Value of TEST_ENV_TEST:</h3>
						<p>` + text + `</p>
					</div>
				</body>
				</html>
				`
	return html
}
