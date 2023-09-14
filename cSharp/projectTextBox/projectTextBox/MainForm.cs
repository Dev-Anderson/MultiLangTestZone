/*
 * Created by SharpDevelop.
 * User: Citeb
 * Date: 14/09/2023
 * Time: 08:50
 * 
 * To change this template use Tools | Options | Coding | Edit Standard Headers.
 */
using System;
using System.Collections.Generic;
using System.Drawing;
using System.Windows.Forms;

namespace projectTextBox
{
	/// <summary>
	/// Description of MainForm.
	/// </summary>
	public partial class MainForm : Form
	{
		public string Anderson; 
		public MainForm()
		{
			//
			// The InitializeComponent() call is required for Windows Forms designer support.
			//
			InitializeComponent();
			
			Anderson = "false"; 
			 
			
			//
			// TODO: Add constructor code after the InitializeComponent() call.
			//
		}
		
		
		
		public static class SharedData
		{
			public static string textbox01 { get; set; }
			public static string textbox02 { get; set; }
		}
		
		public static class ValorDigitado
		{
			public static string alterado {get; set;}
			public static string asdf {get; set;}
			public static string txt01 { get; set; }
			public static string txt02 { get; set; }
		}
		
		void Button1Click(object sender, EventArgs e)
		{
			var form1 = new Form1();
			form1.ShowInTaskbar = false; 
			form1.ShowDialog(); 
			form1.Dispose(); 
		}
		void MainFormLoad(object sender, EventArgs e)
		{
			ValorDigitado.alterado = "false";
			ValorDigitado.asdf = "asdf"; 
			SharedData.textbox01 = "xxx";
			SharedData.textbox02 = "yyy";
		}
	}
}
