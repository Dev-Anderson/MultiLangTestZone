/*
 * Created by SharpDevelop.
 * User: Citeb
 * Date: 14/09/2023
 * Time: 09:54
 * 
 * To change this template use Tools | Options | Coding | Edit Standard Headers.
 */
using System;
using System.Drawing;
using System.Windows.Forms;

namespace projectTextBox
{
	/// <summary>
	/// Description of Form3.
	/// </summary>
	public partial class Form3 : Form
	{
		public Form3()
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
			richTextBox.AppendText(MainForm.ValorDigitado.txt01);
			richTextBox.AppendText("\n");
			richTextBox.AppendText(MainForm.ValorDigitado.txt02); 
		}
	}
}
