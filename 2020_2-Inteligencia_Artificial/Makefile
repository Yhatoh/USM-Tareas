texfile=PlantillaEntrega2
bibfile = Referencias.bib

documento: $(texfile).tex $(bibfile)
	latex $(texfile).tex
	bibtex $(texfile).aux
	latex $(texfile).tex
	bibtex $(texfile).aux
	latex $(texfile).tex
	dvipdf $(texfile).dvi

clean:
	rm -f *~
	rm -f $(texfile).aux
	rm -f $(texfile).bbl
	rm -f $(texfile).blg
	rm -f $(texfile).dvi
	rm -f $(texfile).log
	rm -f $(texfile).out
	rm -f $(texfile).spl
	rm -f $(texfile).pdf
	rm -f $(texfile).tex.backup
	rm -f $(bibfile).backup

