/*
 * Created by SharpDevelop.
 * User: Citeb
 * Date: 14/09/2023
 * Time: 09:11
 * 
 * To change this template use Tools | Options | Coding | Edit Standard Headers.
 */
using System;
using System.Drawing;
using System.Windows.Forms;

namespace projectTextBox
{
	/// <summary>
	/// Description of Form1.
	/// </summary>
	public partial class Form1 : Form
	{
		private string estado; 
		private string valorTxt01; 
		public Form1()
		{
			//
			// The InitializeComponent() call is required for Windows Forms designer support.
			//
			InitializeComponent();
			MainForm mainForm = new MainForm(); 
			estado = mainForm.Anderson; 
			

			//
			// TODO: Add constructor code after the InitializeComponent() call.
			//
		}
		void Label1Click(object sender, EventArgs e)
		{
	
		}
		void Button1Click(object sender, EventArgs e)
		{
			 
			var mainForm = new MainForm();
			mainForm.ShowInTaskbar = false; 
			mainForm.ShowDialog(); 
		}
		void Button2Click(object sender, EventArgs e)
		{
			MainForm.ValorDigitado.txt01 = textBox1.Text; 
			
			var form2 = new Form2();
			form2.ShowInTaskbar = false; 
			form2.ShowDialog(); 
		}

		void TextBox1TextChanged(object sender, EventArgs e)
		{
			estado = "true"; 
		}
		void Form1Load(object sender, EventArgs e)
		{
			valorTxt01 = textBox1.Text;
		}
		void BtnVerClick(object sender, EventArgs e)
		{
			MessageBox.Show(estado);
			MessageBox.Show(textBox1.Text); 
		}
	}
}
