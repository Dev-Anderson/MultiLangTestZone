/*
 * Created by SharpDevelop.
 * User: Citeb
 * Date: 14/09/2023
 * Time: 09:13
 * 
 * To change this template use Tools | Options | Coding | Edit Standard Headers.
 */
using System;
using System.Drawing;
using System.Windows.Forms;

namespace projectTextBox
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
		void BtnVoltarClick(object sender, EventArgs e)
		{
						
			
			var form1 = new Form1(); 
			form1.ShowInTaskbar = false; 
			form1.Visible = true; 
		}
		void Form2Load(object sender, EventArgs e)
		{
//			textBox2.Text = MainForm.SharedData.textbox02;
			textBox2.Text = MainForm.ValorDigitado.alterado; 
			
		}
		void BtnsairClick(object sender, EventArgs e)
		{
			MainForm.ValorDigitado.txt02 = textBox2.Text; 
			
			var form3 = new Form3();
			form3.ShowInTaskbar = false; 
			form3.ShowDialog(); 
		}
	}
}
