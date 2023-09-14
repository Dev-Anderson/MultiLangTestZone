/*
 * Created by SharpDevelop.
 * User: Citeb
 * Date: 12/09/2023
 * Time: 15:38
 * 
 * To change this template use Tools | Options | Coding | Edit Standard Headers.
 */
namespace TesteForms
{
	partial class MainForm
	{
		/// <summary>
		/// Designer variable used to keep track of non-visual components.
		/// </summary>
		private System.ComponentModel.IContainer components = null;
		private System.Windows.Forms.Button btnOpenForm;
		
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
			this.btnOpenForm = new System.Windows.Forms.Button();
			this.SuspendLayout();
			// 
			// btnOpenForm
			// 
			this.btnOpenForm.Location = new System.Drawing.Point(104, 116);
			this.btnOpenForm.Name = "btnOpenForm";
			this.btnOpenForm.Size = new System.Drawing.Size(207, 65);
			this.btnOpenForm.TabIndex = 0;
			this.btnOpenForm.Text = "Abre Form";
			this.btnOpenForm.UseVisualStyleBackColor = true;
			this.btnOpenForm.Click += new System.EventHandler(this.BtnOpenFormClick);
			// 
			// MainForm
			// 
			this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
			this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
			this.ClientSize = new System.Drawing.Size(429, 322);
			this.Controls.Add(this.btnOpenForm);
			this.Name = "MainForm";
			this.StartPosition = System.Windows.Forms.FormStartPosition.CenterScreen;
			this.Text = "Form Principal";
			this.ResumeLayout(false);

		}
	}
}
