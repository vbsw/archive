/**
 * URL Saver (a tool to manage URLs by keywords)
 * Copyright 2016 Vitali Baumtrok
 * 
 * This file is part of URL Saver.
 * 
 * URL Saver is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * 
 * URL Saver is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * 
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */


package com.gitlab.vbg.urlsaver;


import javafx.beans.property.ReadOnlyObjectProperty;
import javafx.beans.property.StringProperty;
import javafx.beans.value.ChangeListener;
import javafx.beans.value.ObservableValue;
import javafx.collections.ObservableList;
import javafx.event.ActionEvent;
import javafx.event.Event;
import javafx.event.EventHandler;
import javafx.scene.Node;
import javafx.scene.Parent;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.control.Tab;
import javafx.scene.control.TabPane;
import javafx.scene.control.TableColumn;
import javafx.scene.control.TableView;
import javafx.scene.control.TableView.TableViewSelectionModel;
import javafx.scene.control.TextField;
import javafx.scene.control.cell.PropertyValueFactory;
import javafx.stage.Stage;
import javafx.stage.WindowEvent;


/**
 * @author Vitali Baumtrok
 */
final class UI
{
	private static final String FILE_COLUMN_TITLE = "File"; //$NON-NLS-1$
	private static final String FILE_COLUMN_ID = "filecolumn"; //$NON-NLS-1$
	private static final String FILE_COLUMN_FIELD = "name"; //$NON-NLS-1$

	private static final String STATE_COLUMN_TITLE = "Exists"; //$NON-NLS-1$
	private static final String STATE_COLUMN_ID = "statecolumn"; //$NON-NLS-1$
	private static final String STATE_COLUMN_FIELD = "state"; //$NON-NLS-1$

	private static final String LOADED_COLUMN_TITLE = "Loaded (%)"; //$NON-NLS-1$
	private static final String LOADED_COLUMN_ID = "loadedcolumn"; //$NON-NLS-1$
	private static final String LOADED_COLUMN_FIELD = "loaded"; //$NON-NLS-1$

	private static final String EDITED_COLUMN_TITLE = "Edited"; //$NON-NLS-1$
	private static final String EDITED_COLUMN_ID = "editedcolumn"; //$NON-NLS-1$
	private static final String EDITED_COLUMN_FIELD = "edited"; //$NON-NLS-1$

	private static final String URL_COLUMN_TITLE = "URL"; //$NON-NLS-1$
	private static final String URL_COLUMN_ID = "urlcolumn"; //$NON-NLS-1$
	private static final String URL_COLUMN_FIELD = "url"; //$NON-NLS-1$

	private static final String FILE_TABLE_VIEW_ID = "#filetableview"; //$NON-NLS-1$
	private static final String URL_TABLE_VIEW_ID = "#urltableview"; //$NON-NLS-1$
	private static final String TAB_PANE_ID = "#tabpane"; //$NON-NLS-1$
	private static final String FILES_TAB_ID_NAME = "filestab"; //$NON-NLS-1$
	private static final String URLS_TAB_ID_NAME = "urlstab"; //$NON-NLS-1$
	private static final String ABOUT_TAB_ID_NAME = "abouttab"; //$NON-NLS-1$
	private static final String FILE_TEXT_FIELD_ID = "#filetext"; //$NON-NLS-1$
	private static final String CREATE_FILE_BTN_ID = "#createfilebtn"; //$NON-NLS-1$
	private static final String DELETE_FILE_BTN_ID = "#deletefilebtn"; //$NON-NLS-1$
	private static final String RELOAD_BTN_ID = "#reloadbtn"; //$NON-NLS-1$
	private static final String SEARCH_RESULT_LABEL_ID = "#searchresultlabel"; //$NON-NLS-1$
	private static final String URL_TEXT_FIELD_ID = "#urltext"; //$NON-NLS-1$
	private static final String OPEN_IN_BROWSER_BTN_ID = "#openinbrowserbtn"; //$NON-NLS-1$
	private static final String QUIT_BTN_ID = "#quitbtn"; //$NON-NLS-1$

	private static final UI.CloseStageListener closeStageListener = new UI.CloseStageListener ( );
	private static final UI.FileTableSelectionListener fileTableSelectionListener = new UI.FileTableSelectionListener ( );
	private static final UI.URLTableSelectionListener urlTableSelectionListener = new UI.URLTableSelectionListener ( );
	private static final UI.CreateFileButtonListener createFileButtonListener = new UI.CreateFileButtonListener ( );
	private static final UI.DeleteFileButtonListener deleteFileButtonListener = new UI.DeleteFileButtonListener ( );
	private static final UI.ReloadButtonListener reloadButtonListener = new UI.ReloadButtonListener ( );
	private static final UI.URLTextFieldListener urlTextFieldListener = new UI.URLTextFieldListener ( );
	private static final UI.OpenInBrowserButtonListener openInBrowserButtonListener = new UI.OpenInBrowserButtonListener ( );
	private static final UI.QuitButtonListener quitButtonListener = new UI.QuitButtonListener ( );
	private static final UI.URLTabSelectedListener urlTabSelectedListener = new UI.URLTabSelectedListener ( );

	private static double decorationWidth = 0;
	private static double decorationHeight = 0;

	static Stage stage;
	static Scene scene;
	static TabPane tabPane;
	static Tab filesTab;
	static Tab urlsTab;
	static Tab aboutTab;
	static Button createFileButton;
	static Button deleteFileButton;
	static Button reloadButton;
	static TextField fileTextField;
	static TableView<FileTableDAO> fileTable;
	static TableView<URLTableDAO> urlTable;
	static Label searchResultLabel;
	static TextField urlTextField;
	static Button openInBrowserButton;
	static Button quitButton;

	static void init ( final Stage stage )
	{
		final Parent root = Ressource.loadFXML ( );

		UI.stage = stage;
		UI.scene = new Scene ( root, Settings.window.getWidth ( ), Settings.window.getHeight ( ) );

		recreateUI ( root );
		reloadCSS ( );
		selectFirstUrlsItemInFileTable ( );

		UI.stage.setOnCloseRequest ( closeStageListener );
		UI.stage.setScene ( UI.scene );
		UI.stage.setTitle ( Settings.window.getTitle ( ) );
		UI.stage.setMaximized ( Settings.window.isMaximized ( ) );
		UI.stage.show ( );

		decorationWidth = UI.stage.getWidth ( ) - UI.scene.getWidth ( );
		decorationHeight = UI.stage.getHeight ( ) - UI.scene.getHeight ( );
	}

	static void reloadFXML ( )
	{
		final Parent root = Ressource.loadFXML ( );

		recreateUI ( root );
	}

	static void applySettings ( )
	{
		stage.setWidth ( Settings.window.getWidth ( ) + decorationWidth );
		stage.setHeight ( Settings.window.getHeight ( ) + decorationHeight );
		stage.setTitle ( Settings.window.getTitle ( ) );
		stage.setMaximized ( Settings.window.isMaximized ( ) );
	}

	static void reloadCSS ( )
	{
		final String cssURL = Ressource.loadCSSURL ( );
		final ObservableList<String> stylesheets = scene.getStylesheets ( );

		stylesheets.clear ( );
		stylesheets.add ( cssURL );
	}

	static void selectFirstUrlsItemInFileTable ( )
	{
		final int firstUrlsItemIndex = getFirstFileTableUrlsItemIndex ( );

		if ( firstUrlsItemIndex >= 0 )
		{
			final TableViewSelectionModel<FileTableDAO> selectionModel = fileTable.getSelectionModel ( );

			selectionModel.select ( firstUrlsItemIndex );
		}
	}

	static void refreshFileItem ( final FileTableDAO item, final boolean selectFocus )
	{
		final ObservableList<FileTableDAO> items = fileTable.getItems ( );
		final TableViewSelectionModel<FileTableDAO> selectionModel = fileTable.getSelectionModel ( );
		final int index = getIndexOfItem ( item );

		items.set ( index, item );

		if ( selectFocus )
		{
			selectionModel.focus ( index );
			selectionModel.select ( index );
			fileTable.requestFocus ( );
		}
		else
		{
			final FileTableDAO focussedItem = FileTable.selectedItem;
			final int focussedItemIndex = getIndexOfItem ( focussedItem );

			selectionModel.focus ( focussedItemIndex );
		}
	}

	static void refreshURLs ( )
	{
		searchResultLabel.setText ( "0/0 (" + URLTable.file.getNameWithoutExtension ( ) + ")" );
	}

	static void refreshTabs ( )
	{
		final FileTableDAO selectedFile = FileTable.selectedItem;
		final String loaded = selectedFile.getLoaded ( );
		final boolean loadingIncomplete = loaded.equals ( "100" ) == false;
		final boolean noUrlsExtension = selectedFile.isUrlsExtension ( ) == false;

		urlsTab.setDisable ( noUrlsExtension || loadingIncomplete );
	}

	static void refreshFileDetails ( )
	{
		final FileTableDAO selectedFile = FileTable.selectedItem;

		UI.fileTextField.setText ( selectedFile.getName ( ) );
		UI.createFileButton.setDisable ( selectedFile.isAvailable ( ) );
		UI.deleteFileButton.setDisable ( selectedFile.isAvailable ( ) == false );
		UI.reloadButton.setDisable ( selectedFile.isAvailable ( ) == false );
	}

	private static int getIndexOfItem ( final FileTableDAO item )
	{
		final ObservableList<FileTableDAO> items = fileTable.getItems ( );

		for ( int i = 0; i < items.size ( ); i += 1 )
		{
			final FileTableDAO currItem = items.get ( i );

			if ( currItem == item )
			{
				return i;
			}
		}

		return -1;
	}

	private static int getFirstFileTableUrlsItemIndex ( )
	{
		final ObservableList<FileTableDAO> items = fileTable.getItems ( );
		int firstURLsItemIndex = -1;

		for ( int i = 0; i < items.size ( ); i += 1 )
		{
			final FileTableDAO item = items.get ( i );
			final String fileName = item.getName ( );

			if ( UrlsFile.isValidExtension ( fileName ) )
			{
				if ( fileName.equals ( Ressource.EXTERNAL_URLS_FILE ) )
				{
					return i;
				}
				else if ( firstURLsItemIndex < 0 )
				{
					firstURLsItemIndex = i;
				}
			}
		}

		return firstURLsItemIndex;
	}

	private static void recreateUI ( final Parent root )
	{
		tabPane = (TabPane) root.lookup ( TAB_PANE_ID );
		filesTab = createTab ( tabPane, FILES_TAB_ID_NAME, null );
		urlsTab = createTab ( tabPane, URLS_TAB_ID_NAME, urlTabSelectedListener );
		aboutTab = createTab ( tabPane, ABOUT_TAB_ID_NAME, null );
		createFileButton = recreateButton ( root.lookup ( CREATE_FILE_BTN_ID ), createFileButtonListener );
		deleteFileButton = recreateButton ( root.lookup ( DELETE_FILE_BTN_ID ), deleteFileButtonListener );
		reloadButton = recreateButton ( root.lookup ( RELOAD_BTN_ID ), reloadButtonListener );
		fileTextField = (TextField) root.lookup ( FILE_TEXT_FIELD_ID );
		fileTable = recreateFileTable ( root.lookup ( FILE_TABLE_VIEW_ID ) );
		urlTable = recreateURLTable ( root.lookup ( URL_TABLE_VIEW_ID ) );
		searchResultLabel = (Label) root.lookup ( SEARCH_RESULT_LABEL_ID );
		urlTextField = recreateURLTextField ( root.lookup ( URL_TEXT_FIELD_ID ) );
		openInBrowserButton = recreateButton ( root.lookup ( OPEN_IN_BROWSER_BTN_ID ), openInBrowserButtonListener );
		quitButton = recreateButton ( root.lookup ( QUIT_BTN_ID ), quitButtonListener );

		scene.setRoot ( root );
	}

	private static TextField recreateURLTextField ( final Node urlTextFieldNode )
	{
		if ( urlTextFieldNode instanceof TextField )
		{
			final TextField urlTextField = (TextField) urlTextFieldNode;
			final StringProperty textProperty = urlTextField.textProperty ( );

			urlTextField.setOnAction ( urlTextFieldListener );
			textProperty.addListener ( urlTextFieldListener );

			return urlTextField;
		}

		return null;
	}

	private static Button recreateButton ( final Node buttonNode, final EventHandler<ActionEvent> listener )
	{
		if ( buttonNode instanceof Button )
		{
			final Button button = (Button) buttonNode;

			button.setOnAction ( listener );

			return button;
		}

		return null;
	}

	private static TableView<FileTableDAO> recreateFileTable ( final Node fileTableNode )
	{
		if ( fileTableNode instanceof TableView )
		{
			@SuppressWarnings ( "unchecked" )
			final TableView<FileTableDAO> fileTable = (TableView<FileTableDAO>) fileTableNode;
			final ObservableList<FileTableDAO> tableItems = fileTable.getItems ( );
			final TableViewSelectionModel<FileTableDAO> selectionModel = fileTable.getSelectionModel ( );
			final ReadOnlyObjectProperty<FileTableDAO> selectedItemProperty = selectionModel.selectedItemProperty ( );

			addColumnToFileTable ( fileTable, FILE_COLUMN_TITLE, FILE_COLUMN_ID, FILE_COLUMN_FIELD );
			addColumnToFileTable ( fileTable, STATE_COLUMN_TITLE, STATE_COLUMN_ID, STATE_COLUMN_FIELD );
			addColumnToFileTable ( fileTable, LOADED_COLUMN_TITLE, LOADED_COLUMN_ID, LOADED_COLUMN_FIELD );

			fileTable.setColumnResizePolicy ( TableView.CONSTRAINED_RESIZE_POLICY );
			tableItems.addAll ( FileTable.urlsFileList );
			tableItems.addAll ( FileTable.propertyDAO, FileTable.fxmlDAO, FileTable.cssDAO );
			selectedItemProperty.addListener ( fileTableSelectionListener );

			return fileTable;
		}

		return null;
	}

	private static TableView<URLTableDAO> recreateURLTable ( final Node fileTableNode )
	{
		if ( fileTableNode instanceof TableView )
		{
			@SuppressWarnings ( "unchecked" )
			final TableView<URLTableDAO> fileTable = (TableView<URLTableDAO>) fileTableNode;
			final TableViewSelectionModel<URLTableDAO> selectionModel = fileTable.getSelectionModel ( );
			final ReadOnlyObjectProperty<URLTableDAO> selectedItemProperty = selectionModel.selectedItemProperty ( );

			addColumnToURLTable ( fileTable, EDITED_COLUMN_TITLE, EDITED_COLUMN_ID, EDITED_COLUMN_FIELD, 70 );
			addColumnToURLTable ( fileTable, URL_COLUMN_TITLE, URL_COLUMN_ID, URL_COLUMN_FIELD, -1 );

			fileTable.setColumnResizePolicy ( TableView.CONSTRAINED_RESIZE_POLICY );
			selectedItemProperty.addListener ( urlTableSelectionListener );

			return fileTable;
		}

		return null;
	}

	private static void addColumnToFileTable ( final TableView<FileTableDAO> fileTable, final String title, final String id, final String daoPropName )
	{
		final TableColumn<FileTableDAO, String> tableColumn = new TableColumn<FileTableDAO, String> ( title );
		final ObservableList<TableColumn<FileTableDAO, ?>> columns = fileTable.getColumns ( );

		tableColumn.setId ( id );
		tableColumn.setCellValueFactory ( new PropertyValueFactory<FileTableDAO, String> ( daoPropName ) );
		columns.add ( tableColumn );
	}

	private static void addColumnToURLTable ( final TableView<URLTableDAO> fileTable, final String title, final String id, final String daoPropName, final int width )
	{
		final TableColumn<URLTableDAO, String> tableColumn = new TableColumn<URLTableDAO, String> ( title );
		final ObservableList<TableColumn<URLTableDAO, ?>> columns = fileTable.getColumns ( );

		tableColumn.setId ( id );
		tableColumn.setCellValueFactory ( new PropertyValueFactory<URLTableDAO, String> ( daoPropName ) );

		if ( width > 0 )
		{
			tableColumn.setMinWidth ( width );
			tableColumn.setMaxWidth ( width );
		}
		columns.add ( tableColumn );
	}

	private static Tab createTab ( final TabPane tabPane, final String tabId, final EventHandler<Event> listener )
	{
		for ( Tab tab: tabPane.getTabs ( ) )
		{
			final String currTabId = tab.getId ( );

			if ( currTabId.equals ( tabId ) )
			{
				if ( listener != null )
				{
					tab.setOnSelectionChanged ( listener );
				}

				return tab;
			}
		}

		return null;
	}

	private static final class CloseStageListener implements EventHandler<WindowEvent>
	{
		@Override
		public void handle ( final WindowEvent event )
		{
			UIController.handeCloseWindow ( event );
		}
	}

	private static final class CreateFileButtonListener implements EventHandler<ActionEvent>
	{
		@Override
		public void handle ( final ActionEvent event )
		{
			UIController.handleCreateSelectedFile ( );
		}
	}

	private static final class DeleteFileButtonListener implements EventHandler<ActionEvent>
	{
		@Override
		public void handle ( final ActionEvent event )
		{
			UIController.handleDeleteSelectedFile ( );
		}
	}

	private static class URLTextFieldListener implements ChangeListener<String>, EventHandler<ActionEvent>
	{
		@Override
		public void changed ( final ObservableValue<? extends String> observable, final String oldValue, final String newValue )
		{
			UIController.handleURLTextChange ( );
		}

		@Override
		public void handle ( final ActionEvent event )
		{
			UIController.handleOpenInBroser ( );
		}
	}

	private static final class OpenInBrowserButtonListener implements EventHandler<ActionEvent>
	{
		@Override
		public void handle ( final ActionEvent event )
		{
			UIController.handleOpenInBroser ( );
		}
	}

	private static final class ReloadButtonListener implements EventHandler<ActionEvent>
	{
		@Override
		public void handle ( final ActionEvent event )
		{
			UIController.handeReloadFile ( );
		}
	}

	private final static class FileTableSelectionListener implements ChangeListener<FileTableDAO>
	{
		@Override
		public void changed ( final ObservableValue<? extends FileTableDAO> observable, final FileTableDAO oldValue, final FileTableDAO newValue )
		{
			if ( newValue != null )
			{
				UIController.handleSelectedFileItem ( newValue );
			}
		}
	}

	private final static class URLTableSelectionListener implements ChangeListener<URLTableDAO>
	{
		@Override
		public void changed ( final ObservableValue<? extends URLTableDAO> observable, final URLTableDAO oldValue, final URLTableDAO newValue )
		{
			if ( newValue != null )
			{
				UIController.handleSelectedURLItem ( newValue );
			}
		}
	}

	private static final class QuitButtonListener implements EventHandler<ActionEvent>
	{
		@Override
		public void handle ( final ActionEvent event )
		{
			UIController.handleQuitApplication ( );
		}
	}

	private static final class URLTabSelectedListener implements EventHandler<Event>
	{
		@Override
		public void handle ( final Event event )
		{
			final Tab tab = (Tab) event.getSource ( );

			if ( tab.isSelected ( ) )
			{
				UIController.handleSelectedURLTab ( );
			}
		}
	}
}
