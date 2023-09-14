/*
 * Created by SharpDevelop.
 * User: Citeb
 * Date: 12/09/2023
 * Time: 15:40
 * 
 * To change this template use Tools | Options | Coding | Edit Standard Headers.
 */
using System;
using System.Drawing;
using System.Windows.Forms;

namespace TesteForms
{
	/// <summary>
	/// Description of Form1.
	/// </summary>
	public partial class Form1 : Form
	{
		public Form1()
		{
			//
			// The InitializeComponent() call is required for Windows Forms designer support.
			//
			InitializeComponent();
			
			//
			// TODO: Add constructor code after the InitializeComponent() call.
			//
		}
		void Button1Click(object sender, EventArgs e)
		{
			var voltarPrincipal = new MainForm(); 
			voltarPrincipal.ShowInTaskbar = false; 
			voltarPrincipal.ShowDialog(); 
		}
		void Button2Click(object sender, EventArgs e)
		{
			var form2 = new Form2(); 
			form2.ShowInTaskbar = false; 
			form2.ShowDialog(); 
		}
	}
}
