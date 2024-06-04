import { Pencil1Icon, TrashIcon } from "@radix-ui/react-icons";
import { useState } from "react";
import EditModal from "../modals/editModal.js";
import styles from '../../shared/style/table.module.css';
import {useDeleteDisciplinesMutation} from "../api/useDeleteDisciplinesMutation";

export function EditCell(props) {
  const deleteDisciplinesMutation = useDeleteDisciplinesMutation();
  const id = props.props.original.id;


  const [isEditModalOpen, setIsEditModalOpen] = useState(false);

  const toggleEditModal = () => {
    setIsEditModalOpen(!isEditModalOpen);
  };
  function deleteDisciplines() {
    if (window.confirm("Вы уверены, что хотите удалить дисциплинe?"))
        deleteDisciplinesMutation.mutateAsync(id);
  }

  return (
    <div style={{padding: '5px'}}>
      <div
        className={styles["btns-wrapper"]}
        onClick={() => toggleEditModal(id)}
      >
        <p>Изменить</p>
        <Pencil1Icon />
      </div>
      <div className={styles["btns-wrapper"]} onClick={deleteDisciplines}>
        <p style={{ color: "var(--warning-color)" }}>Удалить</p>
        <TrashIcon />
      </div>
      {isEditModalOpen && (
        <EditModal
          toggleModal={toggleEditModal}
          isOpen={isEditModalOpen}
          id={id}
        />
      )}
    </div>
  );
}
