<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');

class Home extends CI_Controller {


	/*
		Default	landing page for site.
	*/
	public function index()
	{
		$data['title'] = 'Info';

		$this->load->view('templates/header', $data );
		$this->load->view('pages/info');
		$this->load->view('templates/footer');
	}
	public function contact()
	{
		$data['title'] = 'Contact';
		
		$this->load->view('templates/header', $data );
		$this->load->view('pages/contact.php');
		$this->load->view('templates/footer');
	}
}

/* End of file welcome.php */
/* Location: ./application/controllers/welcome.php */
