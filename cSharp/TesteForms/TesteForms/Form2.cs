/*
 * Created by SharpDevelop.
 * User: Citeb
 * Date: 12/09/2023
 * Time: 15:44
 * 
 * To change this template use Tools | Options | Coding | Edit Standard Headers.
 */
using System;
using System.Drawing;
using System.Windows.Forms;

namespace TesteForms
{
	/// <summary>
	/// Description of Form2.
	/// </summary>
	public partial class Form2 : Form
	{
		public Form2()
		{
			//
			// The InitializeComponent() call is required for Windows Forms designer support.
			//
			InitializeComponent();
			
			//
			// TODO: Add constructor code after the InitializeComponent() call.
			//
		}
		void Button2Click(object sender, EventArgs e)
		{
			this.Close(); 
		}
		void BtnVoltarClick(object sender, EventArgs e)
		{
			var form1 = new Form1(); 
			form1.ShowInTaskbar = false; 
			form1.ShowDialog(); 
		}
	}
}
