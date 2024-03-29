/* --------------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 * ------------------------------------------------------------------------------------------ */

import * as path from 'path';
import {
	workspace,
	ExtensionContext,
	commands,
	window,
	extensions,
	ConfigurationTarget,
	Uri } from 'vscode';

import {
	LanguageClient,
	LanguageClientOptions,
	ServerOptions,
	TransportKind
} from 'vscode-languageclient';

let client: LanguageClient;

export function activate(ctx: ExtensionContext) {
	let plat = process.platform.toString();
	let arch = process.arch;
	let exte = "";
	switch (process.platform.toString()) {
		case "win32": 
			plat = "windows"
			exte = ".exe"
			break;
		default:
			break;
	}
	switch (process.arch) {
		case "x64":
			arch = "amd64"
			break;
		case "x32":
			arch = "386"
			break;
		default:
			break;
	}
	// The server is implemented in Go
	let serverModuleDebug = ctx.asAbsolutePath('adl-lsp');
	let adlc = ctx.asAbsolutePath('adlc');
	let serverModuleRun = ctx.asAbsolutePath(path.join('dist', 'adl-lsp_'+plat+"_"+arch, 'adl-lsp'+exte));
	// The debug options for the server
	// --inspect=6009: runs the server in Node's Inspector mode so VS Code can attach to the server for debugging
	let debugOptions = { execArgv: ['--nolazy', '--inspect=6009'] };

	// If the extension is launched in debug mode then the debug server options are used
	// Otherwise the run options are used
	let serverOptions: ServerOptions = {
		run: { 
			command: serverModuleRun, 
			transport: TransportKind.stdio,
			options: {
				// detached: true
				// ,
				// env: {"ADLC":adlc}
			},
			args: ["stdio", "--adlc-path", adlc],
		},
		debug: { 
			command: serverModuleDebug, 
			transport: TransportKind.stdio, 
			options: {
				// detached: true
				// ,
				// env: {"ADLC":adlc}
			},
			args: ["stdio", "--adlc-path", adlc]
			// args: ["tcp", "client", "--reconnect"]
		}
		// ,
		// run: {
		// 	command: serverModule,
		// 	transport: {
		// 		kind: TransportKind.socket,
		// 		port: 8888
		// 	}
		// },
		// debug: {
		// 	command: serverModule,
		// 	transport: {
		// 		kind: TransportKind.socket,
		// 		port: 8888
		// 	}
		// 	// transport: TransportKind.stdio, //,
		// 	// options: debugOptions
		// }

	};

	// Options to control the language client
	let clientOptions: LanguageClientOptions = {
		// Register the server for plain text documents
		documentSelector: [
			// { scheme: 'file', language: 'plaintext' },
			{ scheme: 'file', language: 'adl' },
			{ scheme: 'file', language: 'tron' },
			{ scheme: 'file', language: 'tron.mod' },
			{ scheme: 'file', language: 'tron.sum' }
		],
		// uriConverters: {
		// 	// Apply file:/// scheme to all file paths.
		// 	code2Protocol: (uri: Uri): string => (uri.scheme ? uri : uri.with({ scheme: 'file' })).toString(),
		// 	protocol2Code: (uri: string) => Uri.parse(uri),
		// },
		// synchronize: {
		// 	// Notify the server about file changes to '.clientrc files contained in the workspace
		// 	fileEvents: workspace.createFileSystemWatcher('**/.clientrc')
		// }
	};

	// Create the language client and start the client.
	client = new LanguageClient(
		'tron',
		'TRON Language Server',
		serverOptions,
		clientOptions
	);

	// Start the client. This will also launch the server
	let languageServerDisposable = client.start();
	ctx.subscriptions.push(languageServerDisposable);

	ctx.subscriptions.push(commands.registerCommand('tron.languageserver.restart', async () => {
		window.showInformationMessage("stopping..")
		await client.stop();
		window.showInformationMessage("stoped..");
		languageServerDisposable.dispose();
		window.showInformationMessage("disposed..");
		languageServerDisposable = client.start();
		window.showInformationMessage("restarted..");
		ctx.subscriptions.push(languageServerDisposable);
	}));

	ctx.subscriptions.push(commands.registerCommand('tron.includes', async () => {
		commands.executeCommand('workbench.action.openWorkspaceSettings');
		const currentDocument = window.activeTextEditor.document;
		const configuration = workspace.getConfiguration('', currentDocument.uri);
		const currentValue = configuration.get('tron.includes', {});
		await configuration.update('tron.includes', currentValue, ConfigurationTarget.WorkspaceFolder);
		commands.executeCommand('settings.switchToJSON');

		// if (window.activeTextEditor) {
		// 	const currentDocument = window.activeTextEditor.document;
		// 	// 1) Get the configuration for the current document
		// 	const configuration = workspace.getConfiguration('', currentDocument.uri);
		// 	// 2) Get the configuration value
		// 	const currentValue = configuration.get('tron.includes', {});
		// 	const target = workspace.workspaceFolders ? ConfigurationTarget.WorkspaceFolder : ConfigurationTarget.Global;
		// 	await configuration.update('tron.includes', ["path here"], target);
		// }
	}));

	ctx.subscriptions.push(commands.registerCommand('tron.show.commands', () => {
		const extCommands = getExtensionCommands();
		window.showQuickPick(extCommands.map(x => x.title)).then(cmd => {
			const selectedCmd = extCommands.find(x => x.title === cmd);
			if (selectedCmd) {
				commands.executeCommand(selectedCmd.command);
			}
		});
	}));
}

export function getExtensionCommands(): any[] {
	const pkgJSON = extensions.getExtension("wxio.tron").packageJSON;
	if (!pkgJSON.contributes || !pkgJSON.contributes.commands) {
		return;
	}
	const extensionCommands: any[] = extensions.getExtension("wxio.tron").packageJSON.contributes.commands.filter((x: any) => x.command !== 'go.show.commands');
	return extensionCommands;
}

export function deactivate(): Thenable<void> | undefined {
	if (!client) {
		return undefined;
	}
	return client.stop();
}
