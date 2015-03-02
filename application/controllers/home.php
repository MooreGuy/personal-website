<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');

class Home extends CI_Controller {


	/*
		Default	landing page for site.
	*/
	public function info()
	{
		$data['title'] = 'Info';
		$data['activeNav'] = 'Info';

		$this->load->view('templates/header', $data );
		$this->load->view('pages/info');
		$this->load->view('tmplates/footer');
	}
}

/* End of file welcome.php */
/* Location: ./application/controllers/welcome.php */
