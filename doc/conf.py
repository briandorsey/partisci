# -*- coding: utf-8 -*-
import sys, os

project = u'Partisci'
copyright = u'2012, Brian Dorsey'

# The short X.Y version.
version = '0.1'
# The full version, including alpha/beta/rc tags.
release = '0.1'

templates_path = ['_templates']
source_suffix = '.rst'
master_doc = 'index'
extensions = ["sphinxcontrib.programoutput"]
exclude_patterns = ['_build']
pygments_style = 'sphinx'

# -- Options for HTML output ---------------------------------------------------
html_theme = 'default'
html_static_path = ['_static']
htmlhelp_basename = 'Partiscidoc'

# -- Options for LaTeX output --------------------------------------------------

latex_elements = {} 
latex_documents = [
  ('index', 'Partisci.tex', u'Partisci Documentation',
   u'Brian Dorsey', 'manual'),
]

# -- Options for manual page output --------------------------------------------

man_pages = [
    ('index', 'partisci', u'Partisci Documentation',
     [u'Brian Dorsey'], 1)
]

# -- Options for Texinfo output ------------------------------------------------

texinfo_documents = [
  ('index', 'Partisci', u'Partisci Documentation',
   u'Brian Dorsey', 'Partisci', 'One line description of project.',
   'Miscellaneous'),
]
