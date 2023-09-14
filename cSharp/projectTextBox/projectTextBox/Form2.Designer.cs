/*
 * Created by SharpDevelop.
 * User: Citeb
 * Date: 14/09/2023
 * Time: 09:13
 * 
 * To change this template use Tools | Options | Coding | Edit Standard Headers.
 */
namespace projectTextBox
{
	partial class Form2
	{
		/// <summary>
		/// Designer variable used to keep track of non-visual components.
		/// </summary>
		private System.ComponentModel.IContainer components = null;
		private System.Windows.Forms.Button btnVoltar;
		private System.Windows.Forms.Button btnsair;
		private System.Windows.Forms.TextBox textBox2;
		private System.Windows.Forms.Label label1;
		
		/// <summary>
		/// Disposes resources used by the form.
		/// </summary>
		/// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
		protected override void Dispose(bool disposing)
		{
			if (disposing) {
				if (components != null) {
					components.Dispose();
				}
			}
			base.Dispose(disposing);
		}
		
		/// <summary>
		/// This method is required for Windows Forms designer support.
		/// Do not change the method contents inside the source code editor. The Forms designer might
		/// not be able to load this method if it was changed manually.
		/// </summary>
		private void InitializeComponent()
		{
			this.btnVoltar = new System.Windows.Forms.Button();
			this.btnsair = new System.Windows.Forms.Button();
			this.textBox2 = new System.Windows.Forms.TextBox();
			this.label1 = new System.Windows.Forms.Label();
			this.SuspendLayout();
			// 
			// btnVoltar
			// 
			this.btnVoltar.Location = new System.Drawing.Point(12, 276);
			this.btnVoltar.Name = "btnVoltar";
			this.btnVoltar.Size = new System.Drawing.Size(75, 23);
			this.btnVoltar.TabIndex = 0;
			this.btnVoltar.Text = "voltar";
			this.btnVoltar.UseVisualStyleBackColor = true;
			this.btnVoltar.Click += new System.EventHandler(this.BtnVoltarClick);
			// 
			// btnsair
			// 
			this.btnsair.Location = new System.Drawing.Point(347, 276);
			this.btnsair.Name = "btnsair";
			this.btnsair.Size = new System.Drawing.Size(75, 23);
			this.btnsair.TabIndex = 1;
			this.btnsair.Text = "Proximo";
			this.btnsair.UseVisualStyleBackColor = true;
			this.btnsair.Click += new System.EventHandler(this.BtnsairClick);
			// 
			// textBox2
			// 
			this.textBox2.Font = new System.Drawing.Font("Microsoft Sans Serif", 12F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
			this.textBox2.Location = new System.Drawing.Point(21, 81);
			this.textBox2.Name = "textBox2";
			this.textBox2.Size = new System.Drawing.Size(391, 26);
			this.textBox2.TabIndex = 2;
			// 
			// label1
			// 
			this.label1.Font = new System.Drawing.Font("Microsoft Sans Serif", 16F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
			this.label1.ImageAlign = System.Drawing.ContentAlignment.TopCenter;
			this.label1.Location = new System.Drawing.Point(21, 55);
			this.label1.Name = "label1";
			this.label1.Size = new System.Drawing.Size(391, 23);
			this.label1.TabIndex = 4;
			this.label1.Text = "Texto 02";
			this.label1.TextAlign = System.Drawing.ContentAlignment.TopCenter;
			// 
			// Form2
			// 
			this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
			this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
			this.ClientSize = new System.Drawing.Size(434, 311);
			this.Controls.Add(this.label1);
			this.Controls.Add(this.textBox2);
			this.Controls.Add(this.btnsair);
			this.Controls.Add(this.btnVoltar);
			this.Name = "Form2";
			this.StartPosition = System.Windows.Forms.FormStartPosition.CenterScreen;
			this.Text = "Form2";
			this.Load += new System.EventHandler(this.Form2Load);
			this.ResumeLayout(false);
			this.PerformLayout();

		}
	}
}
